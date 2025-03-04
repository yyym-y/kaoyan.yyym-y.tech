<template>
  <div class="app-content">
    <div v-for="pageNum in pageNums" :key="pageNum" ref="pageRefs">
      <vue-pdf-embed
        v-if="pageVisibility[pageNum]"
        :source="doc"
        :page="pageNum"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>

import VuePdfEmbed from 'vue-pdf-embed'

// optional styles
import 'vue-pdf-embed/dist/styles/annotationLayer.css'
import 'vue-pdf-embed/dist/styles/textLayer.css'

import { computed, nextTick, onBeforeUnmount, ref, watch } from 'vue'
import { useVuePdfEmbed } from 'vue-pdf-embed'

const props = defineProps<{noteId:string}>()

// 获取链接
import { useServeURL } from '@/store/network';
import { storeToRefs } from 'pinia'
const serveStore = useServeURL()
const { mergeURL } = storeToRefs(serveStore)
const pdf_url = computed(() => mergeURL.value(`note/pdfContent?pdfPath=${props.noteId}`))

const pageRefs = ref([])
  const pageVisibility = ref({})
  let pageIntersectionObserver: IntersectionObserver

  const { doc } = useVuePdfEmbed({
    source: pdf_url,
  })

  const pageNums = computed(() =>
    doc.value
      ? [...Array(doc.value.numPages + 1).keys()].slice(1)
      : []
  )

  const resetPageIntersectionObserver = () => {
    pageIntersectionObserver?.disconnect()
    pageIntersectionObserver = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const index = pageRefs.value.indexOf(entry.target)
          const pageNum = pageNums.value[index]
          pageVisibility.value[pageNum] = true
        }
      })
    })
    pageRefs.value.forEach((element) => {
      pageIntersectionObserver.observe(element)
    })
  }

  watch(pageNums, (newPageNums) => {
    pageVisibility.value = { [newPageNums[0]]: true }
    nextTick(resetPageIntersectionObserver)
  })

  onBeforeUnmount(() => {
    pageIntersectionObserver?.disconnect()
  })
</script>

<style>
.app-content {
  padding: 24px 16px;
}

.app-content > * {
  margin: 0 auto;
  padding-bottom: 8px;
  height: 100%;
}

.vue-pdf-embed__page {
  box-shadow: 0 2px 8px 4px rgba(0, 0, 0, 0.1);
}

</style>