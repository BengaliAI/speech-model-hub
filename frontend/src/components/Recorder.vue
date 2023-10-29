<template>
  <div>
    <div class="flex justify-center items-center m-3">
      <button v-if="!recorder" @click="record()" v-show="!loading"
        class="bg-gradient-to-r from-cyan-500 to-teal-400 hover:from-teal-400 hover:to-cyan-500 text-white font-semibold rounded-xl px-6 py-2 shadow-md transition duration-300 ease-in-out">
        Talk
      </button>
      <button v-else @click="stop()" v-show="!loading"
        class="bg-gradient-to-r from-pink-400 to-amber-500 hover:from-amber-500 hover:to-pink-400 text-white font-semibold rounded-xl px-6 py-2 shadow-md transition duration-300 ease-in-out">
        Stop
      </button>
      <button v-show="loading" type="button" class="inline-flex items-center  bg-gradient-to-r from-pink-400 to-amber-500
         text-white font-semibold rounded-xl px-6 py-2 shadow-md" disabled>
        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none"
          viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
          </path>
        </svg>
        Processing...
      </button>
    </div>

    <div class="flex flex-row justify-center items-center">
      <audio v-if="newAudio" :src="newAudioURL" controls></audio>
      <audio v-if="!newAudio" controls></audio>
    </div>


  </div>
</template>

<script setup>
import { ref } from 'vue';
let newAudio = ref(null);
let recorder = ref(null);

const emits = defineEmits(["recordStart", "recordStop"]);

// Loading is fully controlled by the parent
const props = defineProps({
  loading: Boolean,
  ready: Boolean
});

const newAudioURL = computed(() => {
  return URL.createObjectURL(newAudio.value);
});

const record = async () => {
  if (props.ready === false) {
    alert("Please select a model");
    return;
  }
  newAudio.value = null;

  const stream = await navigator.mediaDevices.getUserMedia({
    audio: true,
    video: false
  });
  const options = { mimeType: "audio/webm" };
  const recordedChunks = [];
  recorder.value = new MediaRecorder(stream, options);

  recorder.value.addEventListener("dataavailable", e => {
    if (e.data.size > 0) {
      recordedChunks.push(e.data);
    }
  });

  recorder.value.addEventListener("stop", () => {
    newAudio.value = new Blob(recordedChunks, { type: "audio/webm" });
    emits("recordStop", newAudioURL.value, newAudio.value);
  });
  emits("recordStart");
  recorder.value.start();
};

const stop = () => {
  // Called first
  recorder.value.stop(); // Call to the event listener
  recorder.value = null;
};
</script>