<template>
  <div class="p-5">
    <h1>Declarative Rendering</h1>
    <DeclarativeRendering />
    <div class="bg-green-200 w-full h-1"></div>

    <h1 class="mt-5">Method</h1>
    <Method />
    <div class="bg-green-200 w-full h-1"></div>

    <h1 class="mt-5">V-Model</h1>
    <VModel />
    <div class="bg-green-200 w-full h-1"></div>
  
    <h1 class="mt-5">Kondisional </h1>
    <ConditionalComponent />
    <div class="bg-green-200 w-full h-1"></div>

    <h1 class="mt-5">Loop </h1>
    <LoopComponent />
    <div class="bg-green-200 w-full h-1"></div>
  </div>
  

  <ChildComponent />
</template>

<script>
import ChildComponent from './components/ChildComponent.vue';
import DeclarativeRendering from './components/DeclarativeRendering.vue';
import Method from './components/Method.vue';
import VModel from './components/VModel.vue';
import ConditionalComponent from './components/ConditionalComponent.vue';
import LoopComponent from './components/LoopComponent.vue';

export default {
  components: {LoopComponent, ChildComponent, DeclarativeRendering, Method, VModel, ConditionalComponent}
};
</script>
