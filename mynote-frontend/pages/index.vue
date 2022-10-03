<template>
  <v-app id="inspire">
    <v-app-bar app>
      <v-container class="py-0 fill-height">
        <v-toolbar-title>MyNote</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon v-if="!expand" @click="click">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
        <v-text-field
          label="https://xxx..."
          hide-details="false"
          v-model="inputArticleUrl"
          v-if="expand"
        ></v-text-field>
        <v-btn
          depressed
          color="teal"
          v-if="expand"
          @click="registerArticle"
        >
          登録
        </v-btn>
        <v-btn icon v-if="expand" @click="click">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </v-container>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-row>
          <v-col cols="3">
            <v-sheet rounded="lg">
              <v-list>
                <v-list-item
                  v-for="[icon, text] in links"
                  :key="icon"
                  link
                >
                  <v-list-item-icon>
                    <v-icon>{{ icon }}</v-icon>
                  </v-list-item-icon>

                  <v-list-item-content>
                    <v-list-item-title>{{ text }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col cols="9">
            <v-row>
              <v-col cols="12">
                <strong>最近登録したコンテンツ</strong>
              </v-col>

              <v-col cols="4"
                v-for="article in articles" :key="article"
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
        ['mdi-home', 'HOME'],
        ['mdi-format-list-bulleted', 'マイリスト'],
      ],
      articles: [],
      expand: false,
      inputArticleUrl: '',
      failed: false,
    }
  },
  created: function () {
    this.expand = false

    this.$axios.get('/api/articles', {
      headers: {
        "Authorization": this.$auth.getToken('local')
      },
      params: { since_id: 0, article_count: 3 } 
          }).then((response) => {
            var responseData = response.data
            var articles = responseData.articles
            this.articles = articles
          })
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
              var articles = responseData.articles
              this.articles = articles
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
