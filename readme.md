1、用户中心
    用户登陆->鉴权->响应
    注册->响应

2、鉴权
    1. 生成session
        session->cache->db

    2. 销毁session  
        session->cache->db

    3. session是否有效      
        session->expired

3、响应
    httpresponse