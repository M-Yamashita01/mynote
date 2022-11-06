<template>
  <v-app id="inspire">
    <Header />
    <v-main>
      <v-container>
        <v-row>
          <LeftMenu />
          <v-col cols="9">
            <v-row>
              <v-col cols="12">
                <strong>最近登録したコンテンツ</strong>
              </v-col>

              <v-col cols="4"
                v-for="article, index in articles" :key="`article-${index}`"
              >
                <v-card height="200" v-bind:href=article.url target="_blank">
                  <v-img
                    src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
                    height="100px"
                  ></v-img>

                  <v-card-title class="text-truncate">
                    {{ article.title }}
                  </v-card-title>

                  <v-card-subtitle>
                    {{ article.website_name }}
                  </v-card-subtitle>
                  </v-card>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>


<script>
import Vue from 'vue'
const axios = require('axios')

export default {
  data: () => {
    return {
      links: [
        ["mdi-home", "HOME"],
        ["mdi-format-list-bulleted", "マイリスト"],
      ],
      articles: [],
      expand: false,
      inputArticleUrl: ''
    }
  },
  created: function () {
    this.expand = false
  },
  methods:{
    click: function() {
      this.expand == true ? this.expand = false : this.expand = true
    },
    registerArticle() {
      this.$axios.post('/api/article', {
        article_url: this.inputArticleUrl
      },
      {
        "Authorization": this.$auth.getToken('local')
      },).then((response) => {
        this.$axios.get('/api/articles', {
          headers: {
            "Authorization": this.$auth.getToken('local')
          },
          params: { since_id: 0, article_count: 3 } 
            }).then((response) => {
              var responseData = response.data
              this.articles = responseData.articles
            })
            .catch(e => {
              alert("Failed to get article.")
            })
        })
        .catch(e => {
          alert("Failed to register article.")
        })

      this.expand = false
      this.inputArticleUrl = ''
    }
  },
}
</script>
