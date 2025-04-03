<script setup lang="ts">
import axios from 'axios'

const handleSubmit = async (form$: any) => {
  
  const requestData = form$.requestData

  form$.submitting = true
  form$.cancelToken = form$.$vueform.services.axios.CancelToken.source()

  let response

  try {
    response = await axios.post(
      'http://localhost:8080/api/v1/search', requestData, {
        cancelToken: form$.cancelToken.token,
        headers: {
          'Content-Type': 'application/json'
        }
      })
    console.log('success', response.data)
  } catch (error) {
    console.log('error', error);
  } finally {
   form$.submitting = false 
  }
}
</script>

<template>
  <div class="">
    <Vueform size="lg" endpoint="false" @submit="handleSubmit">
      <StaticElement>
        <h2>Trustify</h2>
      </StaticElement>
      <TextElement :attrs="{ autofocus: true }" label="Headline" placeholder="Headline" name="headline" input-type="text" autocomplete="off"/>
      <TextElement name="content" label="Additional info" placeholder="Context" input-type="textarea" autocomplete="off"/>
      <TextElement name="url" label="Source" placeholder="http://example.com" input-type="text" autocomplete="off"/>
      <ButtonElement name="submit-button" :submits="true" full>
        Search
      </ButtonElement>
    </Vueform>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
