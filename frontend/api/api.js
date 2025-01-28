// api/api.js
const API_BASE_URL = "http://localhost:8080";

const apiRequest = async (url, method, body = null, token = null) => {
  const headers = {
    "Content-Type": "application/json",
  };

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const options = {
    method,
    headers,
  };

  if (body) {
    options.body = JSON.stringify(body);
  }
 try{
  const response = await fetch(`${API_BASE_URL}${url}`, options);
   if (!response.ok) {
       const errorBody = await response.json();
         throw new Error(errorBody.error);

    }
   return await response.json();
  }catch(err){
      throw err
  }
};

const shortenURL = async (original_url, token) => {
  return apiRequest("/shorten", "POST", { original_url }, token);
};

const registerUser = async (username, email, password) => {
  return apiRequest("/auth/register", "POST", { username, email, password });
};

const loginUser = async (username, password) => {
  return apiRequest("/auth/login", "POST", { username, password });
};

const getCurrentUser = async (token) => {
  return apiRequest("/me", "GET", null, token);
};

const getBookmarks = async (token) => {
  return apiRequest("/bookmarks", "GET", null, token);
};

const createBookmark = async (url, description, tags, token) => {
  return apiRequest("/bookmarks", "POST", { url, description, tags }, token);
};

const updateBookmark = async (id, url, description, tags, token) => {
  return apiRequest(
    `/bookmarks/${id}`,
    "PATCH",
    { url, description, tags },
    token
  );
};

const deleteBookmark = async (id, token) => {
  return apiRequest(`/bookmarks/${id}`, "DELETE", null, token);
};

const searchBookmarks = async (query, token) => {
  return apiRequest(`/search?q=${query}`, "GET", null, token);
};


export default {
  shortenURL,
  registerUser,
  loginUser,
  getCurrentUser,
  getBookmarks,
  createBookmark,
  updateBookmark,
  deleteBookmark,
    searchBookmarks
};