<script setup lang="ts">
defineProps<{
    memoId: string,
    memoTitle: string,
    memoContent: string,
}>()
</script>

<template>
    <va-card class="memo-item">
        <va-card-title class="memo-title">
            {{ memoTitle }}
        </va-card-title>
        <va-card-content class="memo-content" v-html="memoContent" />
        <va-card-actions class="flex memo-action">
            <va-button @click="moveToMemoDetail">Detail</va-button>
        </va-card-actions>
    </va-card>
</template>

<script lang="ts">
import {defineComponent, type PropType} from 'vue'

interface Memo {
    title: string
    content: string
}

export default defineComponent({
    data() {
        return {
            memoList: [] as Array<Memo>
        }
    },
    computed: {

    },
    created() {
    },
    mounted() {

    },
    methods: {
        moveToMemoDetail: function() {
            this.$router.push(`/memo/${this.memoId}`)
        }
    }
})
</script>

<style scoped>
    .memo-item {
        width: 200px;
        height: 250px;
        white-space: pre-wrap;
        word-wrap: break-word;
    }

    .memo-content {
        height: 120px;
        overflow: auto;
    }

    .memo-action {
        width: 100%;
        display: flex;
        position: absolute;
        bottom: 0;
    }
</style>