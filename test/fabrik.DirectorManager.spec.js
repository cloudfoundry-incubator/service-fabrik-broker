'use strict';

const lib = require('../lib');
const catalog = lib.models.catalog;
const proxyquire = require('proxyquire');
const Promise = require('bluebird');
const errors = require('../lib/errors');
const ServiceInstanceAlreadyExists = errors.ServiceInstanceAlreadyExists;

var used_guid = '4a6e7c34-d97c-4fc0-95e6-7a3bc8030be9';
var free_guid = '87599704-adc9-1acd-0be9-795e6a3bc803';
var boshStub = {
  NetworkSegmentIndex: {
    adjust: function (num) {
      return num;
    },
    findFreeIndex: function () {
      return 2;
    }
  },
  director: {
    getDeploymentNames: function () {
      return Promise.resolve([`service-fabrik-0021-${used_guid}`]);
    },
    getDeploymentNameForInstanceId: function () {
      return Promise.resolve([`service-fabrik-0021-${used_guid}`]);
    }
  }
};

var DirectorManager = proxyquire('../lib/fabrik/DirectorManager', {
  '../bosh': boshStub,
});

describe('fabrik', function () {
  describe('DirectorManager', function () {
    const plan_id = 'bc158c9a-7934-401e-94ab-057082a5073f';
    const xsmall_plan_id = plan_id;
    const small_plan_id = 'bc158c9a-7934-401e-94ab-057082a5073e';
    let killJsHintWarning;
    let manager;

    before(function () {
      manager = new DirectorManager(catalog.getPlan(plan_id));
    });

    describe('#getDeploymentName', function () {
      it('should append guid and network segment index to deployment name', function () {
        expect(manager.plan.id).to.eql(plan_id);
        expect(manager.getDeploymentName(used_guid, '90')).to.eql(`service-fabrik-90-${used_guid}`);
        manager.aquireNetworkSegmentIndex(used_guid)
          .catch(err => expect(err).to.be.instanceof(ServiceInstanceAlreadyExists));
        manager.aquireNetworkSegmentIndex(free_guid).then(index => expect(index).to.eql(2));
      });
    });
    describe('#findNetworkSegmentIndex', function () {
      it('should append guid and network segment index to deployment name', function () {
        manager.findNetworkSegmentIndex(used_guid).then(res => expect(res).to.eql(21));
      });
    });
    describe('#isRestorePossible', function () {
      it('should return false when plan not in restore_predecessors', function () {
        // restore not possible from small to xsmall
        manager = new DirectorManager(catalog.getPlan(xsmall_plan_id));
        manager.update_predecessors = [];
        killJsHintWarning = expect(manager.isRestorePossible(small_plan_id)).to.be.false;
      });
      it('should return true when plan not in restore_predecessors', function () {
        // restore possible from xsmall to small
        manager = new DirectorManager(catalog.getPlan(small_plan_id));
        manager.update_predecessors = [xsmall_plan_id];
        killJsHintWarning = expect(manager.isRestorePossible(xsmall_plan_id)).to.be.true;
      });
    });
    describe('#restorePredecessors', function () {
      it('should return update_predecessors if restore_predecessors is not defined', function () {
        manager = new DirectorManager(catalog.getPlan(small_plan_id));
        manager.update_predecessors = [xsmall_plan_id];
        expect(manager.restorePredecessors).to.eql(manager.update_predecessors);
      });
    });
  });
});