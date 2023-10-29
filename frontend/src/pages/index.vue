<template>
  <div class="bg-slate-100 w-full h-screen">
    <div class="flex items-center justify-center">
      <div class="h-fit w-full p-6 m-2 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-md shadow-md">
        <h1 class="text-3xl font-bold text-white text-center">Shobdo</h1>
        <div class="flex flew-row w-full">
          <p class="text-white text-center text-base m-6 p-2 space-x-10 w-full justify-center">Select your model: </p>
          <select
            class="w-full h-10 px-3 m-6 text-base text-gray-700 placeholder-gray-600 border rounded-lg focus:shadow-outline justify-center items-center self-center"
            v-model="selected">
            <option disabled value="" class="text-sm">Model</option>
            <option v-for="opt in options" :value="opt.value" class="text-sm">
              {{ opt.text }}
            </option>
          </select>
        </div>
        <div class="flex flex-col justify-center items-center">
          <Animation :startAnimation="startAnimation" class="mx-6 my-0 p-6 h-20 w-full"></Animation>
          <Recorder :loading="loadingSpinner" @recordStop="handleBlob" @recordStart="startAnimation = true"
            :ready="selected !== ''" class="p-2 w-full h-auto">
          </Recorder>
        </div>
      </div>


    </div>

    <div class="flex items-center justify-center">
      <div class="p-6 m-2 bg-white rounded-md shadow-md w-full bg-gradient-to-r from-teal-300 to-blue-400">
        <h2 class="text-2xl font-bold text-white text-center">Output Text</h2>
        <textarea class="w-full h-40 px-3 py-2 text-base text-gray-700   rounded-lg focus:shadow-outline"
          v-model="outputText" readonly></textarea>
      </div>
    </div>

  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
import { uploadAudio, getModels } from '@/utils/uploader'

const selected = ref('')
const startAnimation = ref(false)
let loadingSpinner = ref(false)
const outputText = ref('')

const handleBlob = async (blobURL: any, blob: Blob) => {
  startAnimation.value = false
  loadingSpinner.value = true
  const res = await uploadAudio(blob, selected.value)
  loadingSpinner.value = false
  outputText.value = res.transcript
}

interface ModelData {
  data: {
    display_name: string,
    name: string,
  }[]
}

let tmp: ModelData = await getModels()

let options: any[] = []
for (let i = 0; i < tmp.data.length; i++) {
  options.push({ text: tmp.data[i].display_name, value: tmp.data[i].display_name })
}

</script>