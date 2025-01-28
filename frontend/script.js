document.addEventListener("DOMContentLoaded", () => {
    const originalURLInput = document.getElementById("original-url");
 const shortenButton = document.getElementById("shorten-btn");
 const shortURLOutput = document.getElementById("short-url-output");
    const bookmarksContainer = document.getElementById("bookmarks-container");
     const bookmarkFormSection = document.getElementById("bookmark-form-section");
   const authFormSection = document.getElementById("auth-form-section");
  const registerForm = document.getElementById("register-form");
  const loginForm = document.getElementById("login-form");
  const authSection = document.getElementById("auth-section");

  const formURL = document.getElementById("form-url");
  const formDescription = document.getElementById("form-description");
 const formTags = document.getElementById("form-tags");
  const formSaveBtn = document.getElementById("form-save-btn");
  const formCancelBtn = document.getElementById("form-cancel-btn");

const searchBookmarks = document.getElementById("search-bookmarks");
let currentBookmarkId = null;
   let authToken = localStorage.getItem("token");
     const registerUsername = document.getElementById("register-username");
      const registerEmail = document.getElementById("register-email");
       const registerPassword = document.getElementById("register-password");
      const registerButton = document.getElementById("register-btn");
      const cancelRegisterButton = document.getElementById("cancel-register-btn");

       const loginUsername = document.getElementById("login-username");
     const loginPassword = document.getElementById("login-password");
       const loginButton = document.getElementById("login-btn");
       const cancelLoginButton = document.getElementById("cancel-login-btn");


   if(authToken){
          getUser()
          showBookmarks()
   }else {
      showAuthButtons()
   }



     shortenButton.addEventListener("click", async () => {
      const originalURL = originalURLInput.value;
     try {
      const response = await fetch("http://localhost:8080/shorten", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            'Authorization': `Bearer ${authToken}`
          },
          body: JSON.stringify({ original_url: originalURL }),
         });
         const data = await response.json();
          shortURLOutput.textContent = data.short_url;
         } catch (error) {
          shortURLOutput.textContent = "Failed to shorten url"
     }
    });
 bookmarksContainer.addEventListener("click", (event) =>{
     const target = event.target;
       if(target.classList.contains("edit-bookmark-btn")){
            currentBookmarkId = target.dataset.id;
           bookmarkFormSection.style.display = "block";
          formURL.value = target.dataset.url;
        formDescription.value = target.dataset.description;
          formTags.value = target.dataset.tags;
    }else if(target.classList.contains("delete-bookmark-btn")){
            deleteBookmark(target.dataset.id);
      }
 })

  formSaveBtn.addEventListener("click", async () =>{
    const url = formURL.value;
      const description = formDescription.value;
        const tags = formTags.value;
     try {
       let response = null;
         if(currentBookmarkId) {
             response = await fetch(`http://localhost:8080/bookmarks/${currentBookmarkId}`,{
                   method: "PATCH",
                   headers: {
                      "Content-Type": "application/json",
                      'Authorization': `Bearer ${authToken}`
                   },
                 body: JSON.stringify({
                      url: url,
                      description: description,
                     tags: tags.split(",")
                  }),
             })

         } else{
              response = await fetch("http://localhost:8080/bookmarks", {
                method: "POST",
                   headers: {
                     "Content-Type": "application/json",
                     'Authorization': `Bearer ${authToken}`
                  },
                  body: JSON.stringify({
                      url: url,
                      description: description,
                     tags: tags.split(",")
                  }),
             });
         }


          const data = await response.json();
           console.log(data)
        bookmarkFormSection.style.display = "none";
       showBookmarks();
      } catch (error) {
           console.log(error)
         }
        currentBookmarkId = null;
  })
  formCancelBtn.addEventListener("click", () => {
       bookmarkFormSection.style.display = "none";
  })

  authSection.addEventListener("click", (event) =>{
    const target = event.target;
    if(target.id === "show-register-btn"){
        registerForm.style.display = "block";
          loginForm.style.display = "none";
          authFormSection.style.display = "block";
   }else if(target.id === "show-login-btn"){
          loginForm.style.display = "block";
          registerForm.style.display = "none";
         authFormSection.style.display = "block";
     }else if(target.id === "logout-btn"){
         localStorage.removeItem("token");
          authToken = null;
          showAuthButtons();
           bookmarksContainer.innerHTML = "";
      }
})
 cancelRegisterButton.addEventListener("click", () => {
     registerForm.style.display = "none";
 })
  cancelLoginButton.addEventListener("click", () => {
       loginForm.style.display = "none";
  })

  registerButton.addEventListener("click", async () =>{
       const username = registerUsername.value;
      const email = registerEmail.value;
      const password = registerPassword.value;

   try {
      const response = await fetch("http://localhost:8080/auth/register", {
          method: "POST",
           headers: {
               "Content-Type": "application/json",
             },
         body: JSON.stringify({
            username: username,
             email: email,
              password: password,
            }),
     });
      const data = await response.json();
      console.log(data);
     registerForm.style.display = "none";
    } catch (error) {
         console.log(error);
      }
   })

  loginButton.addEventListener("click", async () =>{
       const username = loginUsername.value;
        const password = loginPassword.value;

       try {
          const response = await fetch("http://localhost:8080/auth/login", {
             method: "POST",
             headers: {
                "Content-Type": "application/json",
               },
             body: JSON.stringify({
                   username: username,
                password: password,
              }),
           });
           const data = await response.json();
          console.log(data);
          localStorage.setItem("token", data.token);
           authToken = data.token;
           getUser();
          showBookmarks();
        loginForm.style.display = "none";
         authFormSection.style.display = "none";

        } catch (error) {
          console.log(error)
        }
     })
  searchBookmarks.addEventListener("input", (event) =>{
      search(event.target.value);
  })

  async function search(query) {
     try{
      const response = await fetch(`http://localhost:8080/search?q=${query}`,{
        method: "GET",
         headers:{
            'Authorization': `Bearer ${authToken}`
        }
     })
       const data = await response.json();
     renderBookmarks(data);
     }catch(error){
         console.log(error)
    }

  }

  async function showBookmarks(){
        try{
            const response = await fetch("http://localhost:8080/bookmarks", {
                  method: "GET",
                   headers:{
                      'Authorization': `Bearer ${authToken}`
                  }
                })
              const data = await response.json();
            renderBookmarks(data);

         }catch(error){
          console.log(error)
         }
  }

 function renderBookmarks(bookmarks){
      bookmarksContainer.innerHTML = "";
     bookmarks.forEach(bookmark => {
        const bookmarkItem = document.createElement("div");
      bookmarkItem.classList.add("bookmark-item");
        bookmarkItem.innerHTML = `
           <p>URL: <a href="${bookmark.url}" target="_blank">${bookmark.url}</a></p>
           <p>Description: ${bookmark.description}</p>
           <p>Tags: ${bookmark.tags.join(", ")}</p>
           <button class="button is-small is-primary edit-bookmark-btn"
               data-id="${bookmark.id}"
             data-url="${bookmark.url}"
             data-description="${bookmark.description}"
            data-tags="${bookmark.tags}"
               >Edit</button>
           <button class="button is-small is-danger delete-bookmark-btn" data-id="${bookmark.id}">Delete</button>
         `;
       bookmarksContainer.appendChild(bookmarkItem);
     })
  }
async function deleteBookmark(id) {
      try {
           const response = await fetch(`http://localhost:8080/bookmarks/${id}`,{
               method: "DELETE",
                headers:{
                   'Authorization': `Bearer ${authToken}`
              }
             })
             const data = await response.json()
           console.log(data)
          showBookmarks()
       } catch (error) {
          console.log(error)
       }
}
 async function getUser() {
    try{
       const response = await fetch("http://localhost:8080/me",{
           method: "GET",
              headers:{
                 'Authorization': `Bearer ${authToken}`
              }
       });
      const data = await response.json();
    showAuthButtons(data.username);
   }catch(error){
          localStorage.removeItem("token");
          authToken = null;
          showAuthButtons();
   }

 }
 function showAuthButtons(username) {
   authSection.innerHTML = "";
      if(!username){
        authSection.innerHTML = `
           <button id="show-register-btn" class="button is-small is-primary">Register</button>
            <button id="show-login-btn" class="button is-small is-link">Login</button>
          `
      }else{
        authSection.innerHTML = `
            <span class="mr-2">Welcome ${username}</span>
          <button id="logout-btn" class="button is-small is-danger">Logout</button>
          <button id="add-bookmark-btn" class="button is-small is-primary">Add bookmark</button>
        `
          const addBookmarkButton = document.getElementById("add-bookmark-btn");
          addBookmarkButton.addEventListener("click", () =>{
               currentBookmarkId = null;
                bookmarkFormSection.style.display = "block";
            formURL.value = "";
             formDescription.value = "";
            formTags.value = "";
           })
      }
  }

});