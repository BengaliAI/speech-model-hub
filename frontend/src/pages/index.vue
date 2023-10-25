<template>
  <div>
    <!-- <NuxtWelcome /> -->
    <div class="flex items-center justify-center">

      <div class="h-92 w-4/5 p-6 m-6 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-md shadow-md">
        <h1 class="text-3xl font-bold text-white text-center">Shobdo</h1>
        <div class="flex flew-row w-full">
          <p class="text-white text-center m-6 p-2 space-x-10 w-full justify-center">Select your model: </p>
          <select
            class="w-full h-10 px-3 m-6 text-base text-gray-700 placeholder-gray-600 border rounded-lg focus:shadow-outline justify-center items-center self-center"
            v-model="selected">
            <option disabled value="" class="text-sm">Select Model</option>
            <option v-for="option in options" :value="option.value" class="text-sm">
              {{ option.text }}
            </option>
          </select>
        </div>
        <div class="flex flex-col justify-center items-center">
          <!-- <textarea
            class="w-full h-32 px-3 py-2 text-base text-gray-700 placeholder-gray-600 border rounded-lg focus:shadow-outline m-2"
            placeholder="Enter your text here"></textarea> -->
          <Animation :startAnimation="true" class="m-6 p-6 h-20 w-full"></Animation>
          <button
            class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow ">
            Generate
          </button>
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'

const selected = ref('')

const options = [
  { text: 'Alberta', value: 'alberta' },
  { text: 'Bert', value: 'bert' },
  { text: 'Roberta', value: 'roberta' },
]

</script>