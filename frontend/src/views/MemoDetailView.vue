<template>
    <header>
        <div class="header1">
            <va-button
                round
                size="large"
                icon="arrow_back"
                @click="moveToMemoList"
            />
        </div>
        <div class="header2">
            <div class="va-h1">
                {{ memo.Title }}
            </div>
            <div class="buttons">
                <va-button @click="clickEditButton">Edit</va-button>
                <va-button @click="deleteMemo">Delete</va-button>
            </div>
        </div>
    </header>
    <main>
        <div v-html="memo.Content"></div>
    </main>
    <va-modal
        v-model="isShowEditMemoModal"
        hide-default-actions
    >
        <va-form ref="updateFormRef">
            <div class="va-form-header">
                <va-button
                    round
                    icon="close"
                    style="width: 10px;"
                    @click="isShowEditMemoModal = !isShowEditMemoModal"
                />
            </div>
            <va-input
                v-model="editMemo.Title"
                label="title"
                :rules="[(value) => (value && value.length > 0) || 'title is required']"
            />

            <va-input
                v-model="editMemo.Content"
                type="textarea"
                label="content"
                autosize
                :rules="[(value) => (value && value.length > 0) || 'content is required']"
            />

            <va-button :disalbed="!isValid" @click="editMemo()">
                Submit
            </va-button>
        </va-form>
    </va-modal>
</template>

<script lang="ts">
import {defineComponent, type PropType} from 'vue'
import {colorsPreset, useForm} from 'vuestic-ui'
import MemoApiService from "../services/MemoApiService"

interface Memo {
    id: string,
    title: string
    content: string
}

export default defineComponent({
    data() {
        return {
            memo: {},
            editMemo: {},
            isShowEditMemoModal: false as boolean,
            isValid: useForm('updateFormRef'),
        }
    },
    created() {
        this.getMemo(true)
    },
    methods: {
        getMemo: async function(isMd2Html: boolean) {
            const res = await MemoApiService.get(this.$route.params.id, isMd2Html)
            if(!isMd2Html) {
                this.editMemo = await res.data
            } else {
                this.memo = await res.data
            }
        },
        editMemo: async function() {
            const body = JSON.stringify({
                title: this.editMemoTitle,
                content: this.editMemoContent
            })
            const res = await MemoApiService.update(this.$route.params.id, body)
            if(res.statusText === "No Content") {
                alert("Successfuly edit memo")
                location.reload()
            } else {
                alert("Failed to edit memo")
            }
        },
        deleteMemo: async function() {
            if(confirm("Delete this memo?")) {
                const res = await MemoApiService.delete(this.$route.params.id)
                if(res.statusText === "No Content") {
                    alert("Successfuly delete memo")
                    this.moveToMemoList()
                } else {
                    alert("Failed to delete memo")
                }
            }
        },
        moveToMemoList: function() {
            this.$router.push(`/`)
        },
        clickEditButton: function() {
            this.getMemo(false)
            this.isShowEditMemoModal = !this.isShowEditMemoModal
        },
    }
})
</script>

<style scoped>
  header {
    width: 100%;
    padding: 50px 0;
    display: flex;
    flex-flow: column;
    justify-content: flex-start;

    .header1 {
        padding-bottom: 50px;
    }

    .header2 {
        display: flex;
        justify-content: space-between;

        .buttons {
            display: flex;
            align-items: center;
        }

        .buttons * {
            margin: 0 20px;
        }
    }
  }

  main {
    white-space: pre-wrap;
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