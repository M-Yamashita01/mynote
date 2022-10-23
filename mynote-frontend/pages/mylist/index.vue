<template>
  <v-app id="inspire">
    <Header />
    <v-main>
      <v-container>
        <v-row>
          <v-col cols="3">
            <v-sheet rounded="lg">
              <v-list>
                <v-list-item
                  v-for="link in links"
                  :key="`icon-${link.name}`"
                  :to="link.path"
                  link
                >
                  <v-list-item-icon>
                    <v-icon>{{ link.icon }}</v-icon>
                  </v-list-item-icon>

                  <v-list-item-content>
                    <v-list-item-title>{{ link.name }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col cols="9">
            <v-row>
              <v-col cols="12">
                <strong>マイリスト</strong>
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
        { icon: "mdi-home", name: "HOME", path: "/" },
        { icon: "mdi-format-list-bulleted", name: "マイリスト", path: "/mylist" },
      ],
      articles: [],
      expand: false,
      inputArticleUrl: ''
    }
  },
  methods: {
    getArticles() {
    this.$axios.get('/api/articles', {
      headers: {
        "Authorization": this.$auth.getToken('local')
      },
      params: { since_id: 0, article_count: 12 } 
        }).then((response) => {
          var responseData = response.data
          this.articles = responseData.articles
        })
        .catch(e => {
          alert("Failed to get article.");
          console.log("Failed to get article. err: " + err);
        })
    }
  }
}
</script>
