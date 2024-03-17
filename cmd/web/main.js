const { createApp, ref, onMounted } = Vue

// ============== Vue =============
const app = createApp({
  setup() {
    const tables = ref([])
    const selectTable = ref('')
    const columns = ref([])
    const layoutContent = ref('')

    // ============= 初始化 =============
    onMounted(()=>{
      Get('/api/tables').then(data => {
        tables.value = data
        if (data.length > 0) {
          selectTable.value = data[0]
          setTableColumns(data[0])
        }
      }).catch(e => PopError(e))
    })

    const isColumnDisabled = (column) => {
      return column.column === 'deleted_at' || column.column === 'updated_by' || column.column === 'updated_at' || column.column === 'created_at'
    }
    const setTableColumns = (table) => {
      Get(`/api/tables/${table}/columns`).then(data => {
        columns.value = data
      }).catch(e => PopError(e))
    }

    // ============= 事件操作 =============
    const onTableMenuClick = (table) => {
      if (selectTable.value === table) {
        return
      }
      selectTable.value = table
      setTableColumns(table)
    }

    const onSubmit = () => {
      Post(`/api/tables/${selectTable.value}/generate`, columns.value).then(()=>PopSuccess('生成成功')).catch(e=>PopError(e))
    }

    const onAddColumn = () => {
      columns.value.push({
        column: 'new_column',
        label: 'new_column',
        labelWidth: 100,
        type: 'varchar(255)',
        viewType: 'string',
        placeholder: ''
      })
    }
    const onDeleteColumn = (column) => {
      console.log(column)
    }

    return {
      tables, columns, selectTable, layoutContent,
      onTableMenuClick, isColumnDisabled, onSubmit, onAddColumn, onDeleteColumn
    }
  }
})

// ============== 工具函数封装 ==============

function PopSuccess(msg) {
  ElementPlus.ElMessage({
    message: msg,
    type: 'success',
  })
}
function PopError(e) {
  ElementPlus.ElMessage({
    message: e.toString(),
    type: 'error'
  })
}

async function Get(uri) {
  return fetchData(uri, 'GET')
}
async function Post(uri, data) {
  return fetchData(uri, 'POST', JSON.stringify(data))
}
async function Put(uri, data) {
  return fetchData(uri, 'PUT', JSON.stringify(data))
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


app.use(ElementPlus);
app.mount('#app')
