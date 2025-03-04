import req from "./require";
import result from "@/model/utilModel/result";
import { carousel } from "@/model/bussModel/carousel"
import { vedioType } from "@/model/bussModel/vedioType";
import { vedio } from "@/model/bussModel/vedio";

export const getCarousel = ():Promise<result<carousel[]>> => {
    return req.get('/showCarousel').then(res => {
        return res.data
    })
}

export const getVedioType = ():Promise<result<vedioType[]>> => {
    return req.get('/vedioType').then(res => {
        return res.data
    })
}

export const getIndexVedio = ():Promise<result<Record<number, vedio>>> => {
    return req.get('/showIndex').then(res => {
        return res.data
    })
}