package com.wyun.thrift.server.business;

import com.alibaba.fastjson.JSONObject;

/**
 * User:xsg
 * Date:2017/3/1
 */
public interface IBusinessService {

//    IBusinessService INSTANCE=null;

    JSONObject handle(String operation,JSONObject request);

}
