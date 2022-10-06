<template>
  <v-app id="inspire">
    <v-app-bar app>
      <v-toolbar-title>MyNote</v-toolbar-title>
      <div class="flex-grow-1"></div>

      <v-icon>mdi-plus</v-icon>
      <v-icon>mdi-magnify</v-icon>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-row>
          <v-col cols="3">
            <v-sheet rounded="lg">
              <v-list>
                <v-list-item
                  v-for="[icon, text] in links"
                  :key="`icon-${text}`"
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
      articles: []
    }
  },
  created: function () {
    console.log(this.$auth.getToken('local'))

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
  }
}
</script>
