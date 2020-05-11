/* jshint ignore:start */
'use strict';

const _ = require('lodash');
const config = require('@sf/app-config');
const {
  CONST,
  HttpClient
} = require('@sf/common-utils');

class QuotaClient extends HttpClient {
  constructor(options) {
    super(_.defaultsDeep({
      baseUrl: config.quota_app.quota_app_url,
      headers: {
        Accept: 'application/json'
      },
      followRedirect: false
    }, options));
    this.username = config.quota_app.username;
    this.password = config.quota_app.password;
  }

  async getQuotaValidStatus(options) {
    const orgOrSubaccountId = _.get(options, 'orgOrSubaccountId');
    const res = await this.request({
      url: `${config.quota_app.quota_endpoint}/${orgOrSubaccountId}/quota`,
      method: CONST.HTTP_METHOD.GET,
      auth: {
        user: this.username,
        password: this.password
      },
      qs: _.get(options, 'queryParams'),
      json: true
    }, CONST.HTTP_STATUS_CODE.OK);
    return JSON.parse(res.body).quotaValidStatus;
  }
}

module.exports = QuotaClient;
/* jshint ignore:end */
