import { defineStore } from 'pinia';
import { vedio } from '@/model/bussModel/vedio';

export const useVedio = defineStore('vedio', {
    state : () => ({
        vedioInfo : {} as vedio
    })
})