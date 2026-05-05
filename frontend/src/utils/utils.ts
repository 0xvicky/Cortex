
export const getOwner = (url: string) => {
   
    const parsed = new URL(url)
    const isGithub = parsed.hostname === 'github.com'
    const parts = parsed.pathname.replace(/^\/|\/$/g, '').split('/')
    // if (!isGithub || parts.length < 2) {
    //   setValidRepo(false)
    //   return
    // }
    return {parts, isGithub};
    // const [owner, repo] = parts
}

export const  decodeJobId = (jobId) =>{
    // urlsafe base64 uses - and _ instead of + and /
    const base64 = jobId.replace(/-/g, '+').replace(/_/g, '/');
    console.log( JSON.parse(atob(base64)))
    return JSON.parse(atob(base64));
}