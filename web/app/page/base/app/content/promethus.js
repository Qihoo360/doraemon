import React, { Component } from 'react'
import { Button, Table, message, Popconfirm, Input, Icon } from 'antd'
import Highlighter from 'react-highlight-words'
import { getPromethus, addPromethus, updatePromethus, deletePromethus } from '@apis/promethus'
import CreateEditPromethus from './promethus/create-edit-promethus'

export default class Promethus extends Component {
  state = {
    dataSource: [],
    filterItem: {
      name: false,
      url: false,
    },
  }
  componentDidMount() {
    this.getList()
  }

  getColumnSearchProps = dataIndex => ({
    filterDropdown: ({ setSelectedKeys, selectedKeys, confirm }) => (
      <div style={{ padding: 8 }}>
        <Input
          ref={(node) => {
            this.searchInput = node;
          }}
          placeholder={`Search ${dataIndex}`}
          value={selectedKeys[0]}
          onInput={(e) => { setSelectedKeys(e.target.value ? [e.target.value] : []); this.handleSearch(selectedKeys, confirm, dataIndex) }}
          onBlur={() => this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: false },
          }))}
          style={{ width: 188, marginBottom: 8, display: 'block' }}
        />
      </div>
    ),
    filterIcon: filtered => (
      <Icon type="search"
        onMouseDown={() => {
          this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: true },
          })); setTimeout(() => this.searchInput.focus());
        }}
        style={{ color: filtered ? '#1890ff' : undefined }}
      />
    ),
    onFilter: (value, record) => {
      let content
      content = record[dataIndex]
      return content
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase())
    },
    onFilterDropdownVisibleChange: (visible) => {
      if (visible) {
        setTimeout(() => this.searchInput.focus());
      }
    },
    render: text =>
      (this.state.searchedColumn === dataIndex ? (
        <Highlighter
          highlightStyle={{ backgroundColor: '#ffc069', padding: 0 }}
          searchWords={[this.state.searchText]}
          autoEscape
          textToHighlight={text.toString()}
        />
      ) : (
          text
        )),
  })

  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    this.setState({
      searchText: selectedKeys[0],
      searchedColumn: dataIndex,
    });
  }

  getList() {
    getPromethus({}, (res) => {
      this.setState({
        dataSource: res,
      })
    })
  }

  confirmDelete(record) {
    const { id, name } = record
    deletePromethus({}, { id }, (res) => {
      message.success(`删除 ${name} 成功`)
      this.getList()
    })
  }
  handleAdd = (e) => {
    this.createEditPromethus.updateValue()
  }
  handleEdit(record) {
    this.createEditPromethus.updateValue(record)
  }
  onRefRro(component) {
    this.createEditPromethus = component
  }
  updatePromethus = values => new Promise((resolve) => {
    const { id, ...data } = values
    if (id) {
      // edit
      updatePromethus(data, { id }, (res) => {
        this.getList()
        resolve(true)
      })
      return
    }
    addPromethus(data, (res) => {
      this.getList()
      resolve(true)
    })
  })
  render() {
    const { dataSource } = this.state
    const columns = [
      { title: '编号', align: 'center', dataIndex: 'id', key: 'id' },
      {
        title: '名称',
        align: 'center',
        dataIndex: 'name',
        key: 'name',
        ...this.getColumnSearchProps('name'),
        filterDropdownVisible: this.state.filterItem.name,
      },
      {
        title: 'URL',
        align: 'center',
        dataIndex: 'url',
        key: 'url',
        ...this.getColumnSearchProps('url'),
        filterDropdownVisible: this.state.filterItem.url,
      },
      {
        title: '操作',
        align: 'center',
        key: 'action',
        render: (text, record, index) => (
          <span>
            <a onClick={() => this.handleEdit(text, record)}>编辑</a>
            <Popconfirm
              title="确定要删除吗?"
              onConfirm={() => this.confirmDelete(record)}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">删除</a>
            </Popconfirm>
          </span>
        ),
      },
    ]
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <Table dataSource={dataSource} columns={columns} rowKey="id" />
        <CreateEditPromethus OnRef={c => this.onRefRro(c)} onSubmit={this.updatePromethus} />
      </div>
    )
  }
}
