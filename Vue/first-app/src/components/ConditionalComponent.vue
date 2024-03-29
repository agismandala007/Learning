<template>
    <table>
        <tr>
            <td><input v-model="grade" placeholder="input grade..."></td>
        </tr>
        <tr>
            <td>
                <div class="flex flex-row font-bold p-3">
                    <h3>Nilai : </h3>
                    <h3 v-if="grade >= 90" class="bg-green-500 px-3 rounded-md" >A</h3>
                    <h3 v-else-if="grade >= 80 && grade < 90" class="bg-yellow-500 px-3 rounded-md" >B</h3>
                    <h3 v-else-if="grade >= 60 && grade < 80" class="bg-orange-500 px-3 rounded-md">C</h3>
                    <h3 v-else class="bg-red-500 px-3 rounded-md">D</h3>
                </div>
            </td>


        </tr>
    </table>
</template>

<script>
export default {
    data() {
        return {
            grade: 0
        }
    }
}
</script>