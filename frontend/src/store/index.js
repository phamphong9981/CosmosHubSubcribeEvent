import { createStore, createLogger } from 'vuex'
import table from "./modules/table"

const debug = process.env.NODE_ENV !== 'production'

export default createStore({
  modules: {
    table
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})