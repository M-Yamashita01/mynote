<template>
  <v-app id="inspire">
    <v-app-bar app>

      <v-toolbar-title>MyNote</v-toolbar-title>
      <div class="flex-grow-1"></div>

      <v-btn to="/login"
        text
      >
        ログイン
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-row
          no-gutters
          style="flex-wrap: nowrap;"
        >
          <v-col
            cols="6"
            class="flex-grow-0 flex-shrink-0"
          >
            サンプル
          </v-col>
          <v-col
            cols="6"
            class="flex-grow-0 flex-shrink-0"
          >
            <v-container>
              <v-sheet
                color="white"
                elevation="1"
                height="300"
                rounded
                width="300"
              >
                <v-container>
                  <v-form v-model="valid" @submit.prevent="logIn">
                    <p class="text-center">
                      ログイン
                    </p>

                    <v-text-field
                      v-model="email"
                      :rules="emailRules"
                      label="E-mail"
                      required
                    ></v-text-field>

                    <v-text-field
                      v-model="password"
                      :rules="passwordRules"
                      label="Password"
                      required
                    ></v-text-field>

                    <v-btn block color="primary" type="submit">登録</v-btn>
                  </v-form>
                </v-container>
              </v-sheet>
            </v-container>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data(){
      return{
        valid: true,
        email: '',
        emailRules: [
          v => !!v || 'E-mail is required',
          v => /.+@.+/.test(v) || 'E-mail must be valid',
        ],
        password: '',
        passwordRules: [
          v => !!v || 'Password is required',
        ]
      }
    },
    methods: {
      logIn() {
        this.$auth.loginWith('local',{
          data: {
            email: this.email,
            password: this.password,
          }
        })
      }
    },
  }
</script>