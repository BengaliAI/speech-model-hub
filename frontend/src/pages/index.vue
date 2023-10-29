<template>
  <div>
    <div class="flex items-center justify-center">

      <div class="h-fit w-4/5 p-6 m-6 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-md shadow-md">
        <h1 class="text-3xl font-bold text-white text-center">Shobdo</h1>
        <div class="flex flew-row w-full">
          <p class="text-white text-center m-6 p-2 space-x-10 w-full justify-center">Select your model: </p>
          <select
            class="w-full h-10 px-3 m-6 text-base text-gray-700 placeholder-gray-600 border rounded-lg focus:shadow-outline justify-center items-center self-center"
            v-model="selected">
            <option disabled value="" class="text-sm">Model</option>
            <option v-for="option in options" :value="option.value" class="text-sm">
              {{ option.text }}
            </option>
          </select>
        </div>
        <div class="flex flex-col justify-center items-center">
          <Animation :startAnimation="startAnimation" class="mx-6 my-0 p-6 h-20 w-full"></Animation>
          <Recorder :loading="loadingSpinner" @recordStop="printBlobInfo" @recordStart="startAnimation = true"
            class="p-2 w-full h-auto">
          </Recorder>
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
import { uploadAudio } from '@/utils/uploader'

const selected = ref('')
const startAnimation = ref(false)
let loadingSpinner = ref(false)

const printBlobInfo = async (blobURL: any, blob: Blob) => {
  console.log(blobURL, blob)
  startAnimation.value = false
  loadingSpinner.value = true
  const res = await uploadAudio(blob, "test")
  loadingSpinner.value = false
  console.log(res)
}

const options = [
  { text: 'Alberta', value: 'alberta' },
  { text: 'Bert', value: 'bert' },
  { text: 'Roberta', value: 'roberta' },
]

</script>