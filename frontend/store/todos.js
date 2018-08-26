export const state = () => ({
  tasks: {
    todos: [],
    progresses: [],
    dones: []
  }
})

export const mutations = {

  setTodos(state, res) {
    for(let v of res) {
      switch(v.Location) {
        case 'todo':
          state.tasks.todos.push(v)
          break
        case 'progress':
          state.tasks.progresses.push(v)
          break
        case 'done':
          state.tasks.dones.push(v)
          break
      }
    }
    if(state.tasks.todos.length == 0) {
      state.tasks.todos.push({
        Id: 0,
        Text: '',
        Location: 'todo'
      })
    } else if(state.tasks.progresses.length == 0) {
      state.tasks.progresses.push({
        Id: 0,
        Text: '',
        Location: 'progress'
      })
    } else if(state.tasks.dones.length == 0) {
      state.tasks.dones.push({
        Id: 0,
        Text: '',
        Location: 'done'
      })
    }
  },

  add(state, text) {
    state.tasks.todos.push({
      Id: state.todos.length + 1,
      Text: text,
      Done: false,
      Location: 'todo',
      Created_at: Date.now()
    })
  },

  update(state, todo) {
    todo.done = !todo.done
  }

}

export const actions = {
  getTodos({commit}) {
    const method = 'GET'
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    }

    fetch('http://localhost:9999/tasks',{
      method,
      headers
    }).then((res) => res.json()).then((res) => {
        commit("setTodos", res)
    }).catch(console.error)
  },

  addTodo({commit}, text_todo) {
    const obj = { text: text_todo }
    const method = 'POST'
    const body = JSON.stringify(obj);
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    }

    fetch("http://localhost:9999/tasks", {
        method,
        headers,
        body
    }).then((res) => {
      commit('add', obj.text)
    }).catch((err) => {
      console.log(err)
    })
  }
}
