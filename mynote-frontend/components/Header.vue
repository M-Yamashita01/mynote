<template>
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
</template>

<script>

export default {
  data: () => {
    return {
      expand: false,
      inputArticleUrl: ''
    }
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