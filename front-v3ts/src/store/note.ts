// store/note.js
import { defineStore } from 'pinia';

export const useNote = defineStore('note', {
    state: () => ({
        selectNode: {} as { fileType: string; id: string; title: string },
    }),
    actions: {
        setSelectNode(node: { fileType: string; id: string; title: string }) {
            this.selectNode = node;
        },
    },
});