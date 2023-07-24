const root = document.documentElement
const themebtns = document.querySelectorAll('.theme > button')

themebtns.forEach((btn) => {
    btn.addEventListener('click', handleThemeUpdate)
})

function handleThemeUpdate(e) {
    switch(e.target.value) {
        case 'dark':
            root.style.setProperty('--bg', 'black')
            root.style.setProperty('--bg-text', 'white')
            break
        case 'calm':
            root.style.setProperty('--bg', 'lightblue')
            root.style.setProperty('--bg-text', 'white')
            break
        case 'light':
            root.style.setProperty('--bg', 'white')
            root.style.setProperty('--bg-text', 'black')
            break
    } 
}