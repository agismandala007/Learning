<template>
    <div>
        <ol>
            <li v-for="todo in todos" :key="todos.id">{{ todo.activity }}</li>
        </ol>
    </div>
</template>

<script>
export default {
    data () {
        return{
            todos: [
                {id: 1, activity: "eat"},
                {id: 2, activity: "sleep"},
                {id: 3, activity: "code"},
                {id: 4, activity: "repeat"}
            ]
        }
    }
}
</script>