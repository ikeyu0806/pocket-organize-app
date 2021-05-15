<template>
  <section class="section">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-half">
          <b-field label="">
            <b-button type="is-info" @click="get_articles">記事を取得</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  methods: {
    get_articles() {

    }
  },
  mounted: function() {
    const access_token = localStorage.getItem('pocket_access_token')
      axios.get(`http://localhost:5000/get_articles?access_token=${access_token}`)
    .then(response => {
      console.log('status:', response.status)
      console.log('body:', response.data)
    }).catch(err => {
      console.log('err:', err);
    })
  }
}
</script>
