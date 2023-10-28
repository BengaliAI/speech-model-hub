<template>
  <div>
    <div class="flex justify-center items-center m-3">
      <button v-if="!recorder" @click="record()"
        class="bg-gradient-to-r from-cyan-500 to-teal-400 hover:from-teal-400 hover:to-cyan-500 text-white font-semibold rounded-xl px-6 py-2 shadow-md transition duration-300 ease-in-out">
        Talk
      </button>
      <button v-else @click="stop()"
        class="bg-gradient-to-r from-pink-400 to-amber-500 hover:from-amber-500 hover:to-pink-400 text-white font-semibold rounded-xl px-6 py-2 shadow-md transition duration-300 ease-in-out">
        Stop
      </button>
    </div>

    <div class="flex flex-row justify-center items-center">
      <!-- <label v-if="newAudio" class="black text-lg mx-3">Recorded Audio</label> -->
      <audio v-if="newAudio" :src="newAudioURL" controls>Hello</audio>
    </div>


  </div>
</template>

<script setup>
import { ref } from 'vue';

let loading = ref(false);
let newAudio = ref(null);
let recorder = ref(null);

const emits = defineEmits(["recordStart", "recordStop"]);

const newAudioURL = computed(() => {
  return URL.createObjectURL(newAudio.value);
});

const record = async () => {
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