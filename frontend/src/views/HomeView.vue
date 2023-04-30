<script setup lang="ts">
import MemoList from '../components/MemoList.vue'
</script>

<template>
  <header>
    <va-button @click="showPostMemoModal = !showPostMemoModal">Add</va-button>
  </header>
  <main>
    <MemoList />
  </main>
  <va-modal
    v-model="showPostMemoModal"
    hide-default-actions
  >
    <va-form ref="insertFormRef">
      <div class="va-form-header">
        <va-button
          round
          icon="close"
          style="width: 10px;"
          @click="showPostMemoModal = !showPostMemoModal"
        />
      </div>
      <va-input
        v-model="memoTitle"
        label="title"
        :rules="[(value) => (value && value.length > 0) || 'title is required']"
      />

      <va-input
        v-model="memoContent"
        type="textarea"
        label="content"
        autosize
        :rules="[(value) => (value && value.length > 0) || 'content is required']"
      />

      <va-button :disalbed="!isValid" @click="postMemo()">
        Submit
      </va-button>
    </va-form>
  </va-modal>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {useForm} from 'vuestic-ui'
import MemoApiService from "../services/MemoApiService"

export default defineComponent({
    data() {
      return {
        showPostMemoModal: false as boolean,
        isValid: useForm('insertFormRef'),
        memoTitle: "" as string,
        memoContent: "" as string
      }
    },
    methods: {
      postMemo: async function() {
        const body = JSON.stringify({
          title: this.memoTitle,
          content: this.memoContent
        })
        const res = await MemoApiService.insert(body)
        const alertMessage = (res.statusText === "Created" ? "Successfuly insert memo." : "Failed to insert memo.")
        
        alert(alertMessage)
        
        location.reload()
      }
    }
})
</script>

<style scoped>
  header {
    width: 100%;
    padding: 50px 0;
    display: flex;
    flex-flow: row wrap;
    justify-content: flex-end;
  }

  .va-form {
    display: flex;
    flex-flow: column;
    
    .va-form-header {
      display: flex;
      flex-flow: row;
      justify-content: flex-end;
      margin: 0;
    }

    * {
      margin: 10px 0;
    }
  }
</style>