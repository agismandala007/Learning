<template>
    <div class="p-5">
        <h1 class="header">Declarative Rendering</h1>
        <DeclarativeRendering />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Method</h1>
        <Method />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">V-Model</h1>
        <VModel />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Kondisional </h1>
        <ConditionalComponent />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Loop </h1>
        <LoopComponent />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Child Component</h1>
        <ChildComponent />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Props </h1>
        <PropsComponent text="Hello world" />
        <PropsComponent :text="message" />
        <div class="bg-green-200 w-full h-1"></div>

        <h1 class="mt-5 header">Loop Component</h1>
        <PropsComponent v-for="todo in todos" :key="todo" :text="todo" />
        <div class="bg-green-200 w-full h-1"></div>
    </div>



</template>

<script>
import ChildComponent from './components/ChildComponent.vue';
import DeclarativeRendering from './components/DeclarativeRendering.vue';
import Method from './components/Method.vue';
import VModel from './components/VModel.vue';
import ConditionalComponent from './components/ConditionalComponent.vue';
import LoopComponent from './components/LoopComponent.vue';
import PropsComponent from './components/PropsComponent.vue';

export default {
    components: { LoopComponent, ChildComponent, DeclarativeRendering, Method, VModel, ConditionalComponent, PropsComponent },

    data() {
        return {
            message: 'test Props',
            todos: [
                'eat', 'sleep', 'code', 'repeat'
            ]
        }
    }
};
</script>
