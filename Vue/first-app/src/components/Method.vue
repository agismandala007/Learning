<template>
    <table>
        <tr>
            <td class="px-5">
                <h2>{{ count }}</h2>
            </td>
            <td><button @click="count++" class="bg-green-300 outline-2 rounded-md p-2 font-bold">Add 1</button></td>
            <td><button @click="count--" class="bg-red-300 outline-2 rounded-md p-2 font-bold">Minus 1</button></td>
        </tr>
    </table>
</template>

<script>
export default {
    data() {
        return {
            count: 0
        }
    }
}
</script>