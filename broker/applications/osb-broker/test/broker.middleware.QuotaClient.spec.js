'use strict';

const _ = require('lodash');
const config = require('@sf/app-config');
const {
  CONST
} = require('@sf/common-utils');

const QuotaClient = require('../src/api-controllers/middleware/QuotaClient');

describe('#QuotaClient', () => {
  let requestStub;
  const quotaClient = new QuotaClient({});
  const plan_id = 'bc158c9a-7934-401e-94ab-057082a5073f';
  const subaccount_id = 'a8cbbac8-6a20-42bc-b7db-47c205fccf9a';
  const organization_guid = 'b8cbbac8-6a20-42bc-b7db-47c205fccf9a';
  let reqOptions;
  beforeEach(() => {
    requestStub = sinon.stub(quotaClient, 'request');
    reqOptions = {
      subaccountId: subaccount_id,
      queryParams: {
        planId: plan_id,
        previousPlanId: undefined,
        useAPIServerForConsumedQuotaCheck: false,
        reqMethod: 'PATCH',
        orgId: organization_guid
      }
    };
  });
  afterEach(() => {
    requestStub.restore();
  });
  it('should make appropriate call for instance based quota', async () => {
    requestStub.resolves({body: { quotaValidStatus: 0}});
    const { quotaValid } = await quotaClient.checkQuotaValidity(reqOptions, true);
    expect(requestStub).to.have.been.calledWithExactly({
      url: `${config.quota_app.quota_endpoint}/${subaccount_id}/quota`,
      method: CONST.HTTP_METHOD.GET,
      params: _.get(reqOptions, 'queryParams'),
      responseType: 'json'
      },
      CONST.HTTP_STATUS_CODE.OK
    );
    expect(quotaValid).to.be.equal(0);
  });
  it('should make appropriate call for non-instance based quota', async () => {
    reqOptions.data = {
        service_id: 'dummy_service_id',
        context: {}
    };
    requestStub.resolves({body: { quotaValidStatus: 0}});
    const { quotaValid } = await quotaClient.checkQuotaValidity(reqOptions, false);
    expect(requestStub).to.have.been.calledWithExactly({
      url: `${config.quota_app.quota_endpoint}/${subaccount_id}/quota`,
      method: CONST.HTTP_METHOD.PUT,
      data: _.get(reqOptions, 'data'),
      responseType: 'json'
      },
      CONST.HTTP_STATUS_CODE.OK
    );
    expect(quotaValid).to.be.equal(0);
  });
});