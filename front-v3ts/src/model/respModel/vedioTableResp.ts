import { playInfo } from "../bussModel/playInfo";

export interface vedioTableResp {
	vedioId: number,
	vedioName: string,
	vedioCover: string,
	description: string,
	visable: boolean,
	typeId: number,
    episode: number,
    episodeInfos : playInfo[]
}