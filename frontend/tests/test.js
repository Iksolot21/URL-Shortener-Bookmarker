// tests/tests.js

// Mock fetch function for testing API calls
const mockFetch = (data, status = 200) => {
    return jest.fn(() => Promise.resolve({
      status,
       json: () => Promise.resolve(data)
     }))
    }
      global.fetch = mockFetch({})
   
    describe('Tests for shorten functionality', ()=>{
     test('Shorten function makes correct API call', async () =>{
         const originalURL = "https://www.test.com"
         const shortenBtn = document.createElement("button")
         shortenBtn.id = "shorten-btn"
        document.body.appendChild(shortenBtn)
   
        const input = document.createElement("input");
       input.id = "original-url"
        input.value = originalURL
        document.body.appendChild(input)
   
       const output = document.createElement("div");
       output.id = "short-url-output";
       document.body.appendChild(output)
     
       const mockedResponse = {short_url: "short_url_test"}
       global.fetch = mockFetch(mockedResponse)
   
        require('../script')
       shortenBtn.click()
   
        await new Promise(r => setTimeout(r, 100));
        expect(output.textContent).toBe(mockedResponse.short_url)
       })
   });
   
    describe("Tests for Bookmark Functionality", () =>{
       test('Render bookmarks function renders bookmarks correctly', async ()=>{
           const bookmarksContainer = document.createElement("div");
            bookmarksContainer.id = "bookmarks-container";
            document.body.appendChild(bookmarksContainer);
   
             const mockedBookmarks = [
                 {
                   id: 1,
                  url: "https://short.url",
                    description: "test bookmark",
                      tags: ["tag1", "tag2"],
                   },
                 {
                   id: 2,
                      url: "https://another.short.url",
                    description: "another bookmark",
                   tags: ["tag3"],
                  },
              ]
            global.fetch = mockFetch(mockedBookmarks);
          require('../script')
           await new Promise(r => setTimeout(r, 100));
         expect(bookmarksContainer.children.length).toBe(2);
        expect(bookmarksContainer.children[0].textContent).toContain(mockedBookmarks[0].description);
      })
   });
   
   
     describe("Tests for Auth Functionality", () =>{
         test("Get user function should get user data", async () => {
             localStorage.setItem("token", "testToken")
             const authSection = document.createElement("div")
            authSection.id = "auth-section"
            document.body.appendChild(authSection)
           const mockedUser = {username: "testUser"}
           global.fetch = mockFetch(mockedUser);
             require('../script')
           await new Promise(r => setTimeout(r, 100));
             expect(authSection.textContent).toContain(mockedUser.username)
   
       })
   });