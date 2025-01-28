// store/store.js
const state = {
    user: null,
    token: localStorage.getItem("token") || null,
    bookmarks: [],
    loading: false,
    error: null
  };
  
  const mutations = {
    setUser(user) {
      state.user = user;
    },
    setToken(token) {
      state.token = token;
       localStorage.setItem("token", token)
    },
    setBookmarks(bookmarks) {
      state.bookmarks = bookmarks;
    },
    setLoading(loading) {
      state.loading = loading;
    },
    setError(error){
        state.error = error
    }
  };
  
  const getters = {
    getUser: () => state.user,
    getToken: () => state.token,
    getBookmarks: () => state.bookmarks,
    isLoading: () => state.loading,
      getError: () => state.error
  };
  
  const actions = {
    //actions here
  };
  
  export default {
    state,
    mutations,
    getters,
    actions,
  };