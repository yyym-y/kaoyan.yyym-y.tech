import { type vedio } from '@/model/bussModel/vedio';

export interface carousel {
    carouselId : number,
    cover : string,
    ifShowVedio : boolean,
    vedio : vedio
}