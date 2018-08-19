export const state = () => ({
  list: []
})

export const mutations = () => ({

  add(state, text) {
    state.list.push({
      text: text,
      done: false,
    })
  },

  update(state, todo) {
    todo.done = !todo.done
  }

})
