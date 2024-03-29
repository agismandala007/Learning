<template>
<table>
    <tr>
        <td>Name : {{ name }}</td>
    </tr>
    <tr>
        <td><input v-model="name" placeholder="input name..."></td>
    </tr>
</table>
</template>

<script>
export default {
    data() {
        return {
            name: ""
        }
    }
}
</script>