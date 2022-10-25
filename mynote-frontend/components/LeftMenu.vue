<template>
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
    }
  }
}
</script>