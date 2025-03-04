import req from "./require";
import * as user from '@/model/reqModel/userReq'
import * as userReq from '@/model/respModel/userResp'
import result from '@/model/utilModel/result'

export const login = (data:user.loginReq) : Promise<result<userReq.loginResp>> => {
    return req.post('/login', data).then(res => {
        return res.data
    })
}