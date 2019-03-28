const rp = require('request-promise')

const id = "aw1flt"
const baseOptions = {headers: {'User-Agent': 'Request-Promise'},json: true}
let comments = getAllComments(id)

async function getAllComments(id) {
  let {all, more} = await getComments(id)
  let i,j,batch,items,chunk = 100;
  for (i=0,j=more.length; i<j; i+=chunk) {
      batch = more.slice(i,i+chunk)
      items = await getMoreChildren(id, batch)
      all = all.concat(items)
      // do whatever
  }

  console.log(all.length)
}

async function getComments(id){
  let options = {
    uri: `https://api.reddit.com/comments/${id}`,
  }
  let resData = await rp(Object.assign(options, baseOptions))
  let listing = resData.find(({data})=>!data.dist)
  let comments = listing.data.children.filter(l=>l.kind==="t1")
  let more = listing.data.children.filter(l=>l.kind==="more")
  more = more.reduce((prev, curr)=>{
    return prev.concat(curr.data.children)
  }, [])
  let items = comments.map(({data})=>({author:data.author, body: data.body}))
  return {all:items, more}
}

async function getMoreChildren(id, children){
  let options = {
    uri: `https://api.reddit.com/api/morechildren`,
    qs: {
      api_type: "json",
      link_id: `t3_${id}`,
      children: children.join(',')
    }
  }
  let resData = await rp(Object.assign(options, baseOptions))
  let things = resData.json.data.things;
  let comments = things.filter(l=>l.kind==="t1")
  let more = things.filter(l=>l.kind==="more")
  let items = comments.map(({data})=>({author:data.author, body: data.body}))
  return items;
}
