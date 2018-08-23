<template>
  <div>
  <el-input
    type="textarea"
    autosize
    placeholder="请输入内容"
    v-model="textarea2">
  </el-input>
  <div style="margin: 20px 0;"></div>
  <el-input
    type="textarea"
    :autosize="{ minRows: 2, maxRows: 20}"
    placeholder="请输入内容"
    v-model="textarea3">
  </el-input>
  </div>
</template>

<script>
export default {
  data() {
    return {
      textarea2: '',
      textarea3: ''
    }
  }
}
</script>