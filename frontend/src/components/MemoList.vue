<script setup lang="ts">
import MemoItem from './items/MemoItem.vue'
</script>

<template>
    <div class="memo-list" v-if="memoList.length > 0">
        <MemoItem v-for="memo in memoList"
            :memoId="memo.ID"
            :memoTitle="memo.Title"
            :memoContent="memo.Content">
        </MemoItem>
    </div>
    <div class="memo-list-empty" v-else>
        No memos found.
    </div>
</template>

<script lang="ts">
import {defineComponent, type PropType} from 'vue'
import MemoApiService from "../services/MemoApiService"

interface Memo {
    id: string,
    title: string
    content: string
}

export default defineComponent({
    data() {
        return {
            memoList: [],
            md2html: true
        }
    },
    created() {
        this.getMemoList(true)
    },
    methods: {
        getMemoList: async function(isMd2Html: boolean) {
            const res = await MemoApiService.getAll(isMd2Html);
            this.memoList = await res.data
        }
    }
})
</script>

<style scoped>
.memo-list {
    display: flex;
    flex-wrap: wrap;
    gap: 50px 30px;
}

.memo-list-empty {
    display: flex;
    justify-content: center;
}

</style>