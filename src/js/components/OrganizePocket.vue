<template>
  <section class="section">
    <div class="container">
      <b-field label="表示記事数" :label-position="labelPosition">
        <b-numberinput :value="articleCount"
                       class="articleCount">
        </b-numberinput>
      </b-field>
      <b-field label="タグで検索" :label-position="labelPosition" grouped>
        <b-taginput
          v-model="searchTag"
          ellipsis
          icon="label"
          placeholder="タグを入力してください">
        </b-taginput>
      </b-field>
      <b-button
        class="button is-primary search-button"
        @click="searchArticles">
        検索する
      </b-button>
      <b-field label="チェックした記事にタグを付与" :label-position="labelPosition" grouped>
        <b-taginput
          v-model="updateTags"
          ellipsis
          icon="label"
          placeholder="タグを入力してください">
        </b-taginput>
      </b-field>
      <b-button
        class="button is-primary exec-button"
        v-on:click="addTags">更新する</b-button>
      <b-table :data="articles"
               checkable
               :checked-rows.sync="checkedRows"
               v-model="checkedItemIDs"
               :checkbox-position="checkboxPosition">
          <b-table-column field="item_id"
                  label="記事ID"
                  v-slot="props">
          {{ props.row.item_id }}
        </b-table-column>
        <b-table-column field="resolved_title"
                        label="タイトル"
                        v-slot="props">
          <a v-bind:href="props.row.given_url"
             target="_blank"
             rel=“noopener”>{{ articleTitle(props.row.resolved_title) }}</a>
        </b-table-column>
      </b-table>
    </div>
  </section>
</template>

<style>
  .search-button {
    margin-top: 10px;
    margin-bottom: 20px;
  }
  .exec-button {
    margin-top: 10px;
    margin-bottom: 30px;
  }
  .articleCount {
    width: 30%;
  }
</style>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      articles: [],
      searchTag: [],
      updateTags: [],
      checkedItemIDs: ["1902529220"],
      articleCount: 30,
      checkboxPosition: 'left',
      labelPosition: 'on-border',
      checkedRows: []
    }
  },
  methods:{
    articleTitle: function(articleTitle) {
      let result = articleTitle
      if (result == "") {
        result = "タイトルが登録されていません"
      }
      return result
    },
    searchArticles: function() {
      const access_token = localStorage.getItem('pocket_access_token')
      axios.get(`http://localhost:5000/get_articles?access_token=${access_token}&count=${this.articleCount}&tag=${this.searchTag}`)
      .then(response => {
        console.log('status:', response.status)
        console.log('body:', response.data)
        this.articles = Object.values(response.data.list)
      }).catch(err => {
        console.log('err:', err);
      })
    },
    addTags: function() {
      const access_token = localStorage.getItem('pocket_access_token')
      const item_ids = []
      this.checkedRows.forEach((r) => (item_ids.push(r.item_id)))
      axios.post('http://localhost:5000/add_tags', {
        access_token: access_token,
        // TODOここは配列全体を送れるように
        tag: this.updateTags[0],
        item_ids: item_ids[0]
      })
      .then(response => {
        console.log('status:', response.status)
        console.log('body:', response.data)
      }).catch(err => {
        console.log('err:', err);
      })
    }
  },
  beforeCreate: function() {
    const access_token = localStorage.getItem('pocket_access_token')
      axios.get(`http://localhost:5000/get_articles?access_token=${access_token}&count=30&tag=_untagged_`)
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
