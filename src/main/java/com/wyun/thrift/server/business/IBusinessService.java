package com.wyun.thrift.server.business;

import com.alibaba.fastjson.JSONObject;
import com.wyun.thrift.server.Response;

/**
 * User:xsg
 * Date:2017/3/1
 */
public interface IBusinessService {

    Response handle(String operation, JSONObject request);

}
