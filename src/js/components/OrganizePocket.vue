<template>
  <section class="section">
    <div class="container">
      <b-table :data="articles" :columns="columns"></b-table>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      articles: [],
      columns: [
        {
          field: 'resolved_title',
          label: 'resolved_title',
        },
        {
          field: 'given_url',
          label: 'given_url',
        },
        {
          field: 'time_added',
          label: 'time_added',
        },
        {
          field: 'favorite',
          label: 'favorite',
        }
      ]
    }
  },
  mounted: function() {
    const access_token = localStorage.getItem('pocket_access_token')
      axios.get(`http://localhost:5000/get_articles?access_token=${access_token}`)
    .then(response => {
      console.log('status:', response.status)
      console.log('body:', response.data)
      this.articles = Object.values(response.data.list)
    }).catch(err => {
      console.log('err:', err);
    })
  }
}
</script>
