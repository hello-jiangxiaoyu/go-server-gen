const { createApp, ref, onMounted } = Vue

// ============== Vue =============
const app = createApp({
  setup() {
    const tables = ref([])
    const selectTable = ref('')
    const columns = ref([])
    onMounted(()=>{
      Get('/api/tables').then(data => {
        tables.value = data
        if (data.length > 0) {
          selectTable.value = data[0]
          setTableColumns(data[0])
        }
      }).catch(e => PopError(e))
    })

    const setTableColumns = (table) => {
      Get(`/api/tables/${table}/columns`).then(data => {
        columns.value = data
        console.log("columns: ", data)
      }).catch(e => PopError(e))
    }
    const onTableMenuClick = (table) => {
      if (selectTable.value === table) {
        return
      }
      selectTable.value = table
      setTableColumns(table)
    }

    // ============= 数据表处理 =============
    return {
      selectTable, columns, onTableMenuClick, tables
    }
  }
})

function PopError(e) {
  ElementPlus.ElMessage({
    message: e.toString(),
    type: 'error'
  })
}

// ============== fetch 请求封装 ==============
async function fetchData(uri, method, data) {
  let err = '';
  const response = await fetch(uri, {
    method,
    headers: {'content-type': 'application/json'},
    body: data
  }).then((resp) => resp.json())
    .catch((e) => {
      err = e.toString()
    });

  if (err !== '') {
    return Promise.reject('fetch error');
  } else if (typeof response !== 'object') {
    return Promise.reject('Invalid server response type');
  }

  if (response?.msg) {
    return Promise.reject(response.msg);
  }
  return response;
}

async function Get(uri) {
  return fetchData(uri, 'GET')
}
async function Post(uri) {
  return fetchData(uri, 'POST')
}

app.use(ElementPlus);
app.mount('#app')
