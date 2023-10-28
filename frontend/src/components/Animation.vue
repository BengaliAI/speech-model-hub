<template>
    <div class="flex justify-center items-center">
        <div id="bars" class="flex justify-center items-center" v-show="!startAnimation">
            <div v-for="i in  10 " class="bar" :style="{ animationDuration: '0ms' }">
            </div>
        </div>
        <div id="bars" class="flex sm:hidden justify-center items-center" v-show="startAnimation">
            <div v-for=" i  in  10 " class="bar" :style="{ animationDuration: randomDuration(700, 200) + 'ms' }">
            </div>
        </div>
        <div id="bars" class="hidden md:flex justify-center items-center" v-show="startAnimation">
            <div v-for=" i  in  20 " class="bar" :style="{ animationDuration: randomDuration(700, 200) + 'ms' }">
            </div>
        </div>
        <div id="bars" class="hidden lg:flex justify-center items-center" v-show="startAnimation">
            <div v-for=" i  in  22 " class="bar" :style="{ animationDuration: randomDuration(700, 200) + 'ms' }">
            </div>
        </div>
    </div>
</template>
  
<script  setup>
const props = defineProps({
    startAnimation: Boolean
})

function randomDuration(mean, std) {
    let u = 0;
    let v = 0;
    while (u === 0) u = Math.random();
    while (v === 0) v = Math.random();
    const z = Math.sqrt(-2.0 * Math.log(u)) * Math.cos(2 * Math.PI * v);
    const randomValue = mean + z * std;
    return randomValue;

}
</script>
  
<style scoped>
.bar {
    background: #52467b;
    bottom: 1px;
    height: 3px;
    width: 10px;
    margin: 0px 4px;
    border-radius: 5px;
    animation: sound 0ms 600ms linear infinite alternate;
}

@keyframes sound {
    0% {
        opacity: .35;
        height: 3px;
    }

    100% {
        opacity: 1;
        height: 70px;
    }
}
</style>
  
