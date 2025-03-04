<template>
    <div id="editer"></div>
</template>

<script lang="ts" setup>
import Vditor from 'vditor';
import {  onMounted, onUnmounted, ref, } from 'vue';
import 'vditor/dist/index.css';

const props = defineProps<{
    handleContent:Function,
    handleSave:Function
}>()

let content = ref("")
const editer = ref<Vditor | null>(null);
const getContent = async () => {
    content.value = await props.handleContent()
    editer.value?.setValue(content.value, true);
}

const handleSave = (mode:string = "manual") => {
    let content = editer.value?.getValue() || ""
    props.handleSave(content, mode)
}

const refresh = () => {
    getContent()
}

defineExpose({ refresh })

const handleKeyDown = (event:any) => {
    if (event.ctrlKey && event.key === 's') {
        event.preventDefault();
        console.log('Ctrl + S was pressed!');
        handleSave()
    }
}


// 初始化 Vditor
onMounted(() => {
    editer.value = new Vditor('editer', {
        height : "84vh",
        toolbar: [
            {
            name: 'Save',
            tipPosition: 's',
            tip: '保存',
            className: 'right',
            icon: '<svg t="1739630773025" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4190" width="200" height="200"><path d="M925.248 356.928l-258.176-258.176a64 64 0 0 0-45.248-18.752H144a64 64 0 0 0-64 64v736a64 64 0 0 0 64 64h736a64 64 0 0 0 64-64V402.176a64 64 0 0 0-18.752-45.248zM288 144h192V256H288V144z m448 736H288V736h448v144z m144 0H800V704a32 32 0 0 0-32-32H256a32 32 0 0 0-32 32v176H144v-736H224V288a32 32 0 0 0 32 32h256a32 32 0 0 0 32-32V144h77.824l258.176 258.176V880z" p-id="4191"></path></svg>',
            click () {handleSave()},
        }, 'emoji', 'bold', 'quote', 'list', 'ordered-list', 'check', 'outdent', 'indent', 'code', 'inline-code', 'upload', 'link','table','edit-mode', 'fullscreen','outline','code-theme','content-theme','export'],
        toolbarConfig: {
            pin : true,
        },
        after() {
            getContent()
        },
        input(md) {
            handleSave("auto")
        }
    });
    document.addEventListener('keydown', handleKeyDown);
});

onUnmounted(() => {
    editer.value?.destroy()
    document.removeEventListener('keydown', handleKeyDown);
    console.log('Ctrl + S listener removed.');
})
</script>

<style>
</style>