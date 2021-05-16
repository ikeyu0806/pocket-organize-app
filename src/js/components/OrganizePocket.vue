<template>
  <section class="section">
    <div class="container">
      <b-field label="表示記事数" :label-position="labelPosition">
        <b-numberinput laceholder="10" :min="1" :max="100"></b-numberinput>
      </b-field>
      <b-field label="タグで検索" :label-position="labelPosition" grouped>
        <b-taginput
          :value="['レシピ', '映画']"
          ellipsis
          icon="label"
          placeholder="タグを入力してください">
        </b-taginput>
      </b-field>
      <b-field label="チェックされた記事にタグを付与" :label-position="labelPosition" grouped>
        <b-taginput
          :value="['レシピ', '映画']"
          ellipsis
          icon="label"
          placeholder="タグを入力してください">
        </b-taginput>
      </b-field>
      <b-button class="button is-primary">実行する</b-button>
      <b-table :data="articles"
               checkable
               :checkbox-position="checkboxPosition">
        <b-table-column field="resolved_title"
                        label="タイトル"
                        v-slot="props">
          <a v-bind:href="props.row.given_url"
             target="_blank"
             rel=“noopener”>{{ props.row.resolved_title }}</a>
        </b-table-column>
          <b-table-column field="time_added"
                  label="登録日時"
                  v-slot="props">
          {{ props.row.time_added }}
        </b-table-column>
      </b-table>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      articles: [],
      checkboxPosition: 'left'
    }
  },
  beforeCreate: function() {
    const access_token = localStorage.getItem('pocket_access_token')
      axios.get(`http://localhost:5000/get_articles?access_token=${access_token}&count=10&tag=グルメ`)
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
