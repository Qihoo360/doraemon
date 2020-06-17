import React, { Component } from 'react';
import { Button, Table, message, Popconfirm, Input, Icon } from 'antd';
import Highlighter from 'react-highlight-words'
import { getGroup, addGroup, updateGroup, deleteGroup } from '@apis/group';
import CreateEditGroup from './group/create-edit-group';

export default class Rules extends Component {
  state = {
    dataSource: [],
    filterItem: {
      name: false,
      user: false,
    },
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

  componentDidMount() {
    this.getList()
  }
  getList() {
    getGroup({}, (res) => {
      this.setState({
        dataSource: res,
      });
    });
  }
  handleAdd = (e) => {
    this.createEditGroup.updateValue()
  }
  handleEdit = (text, record) => {
    this.createEditGroup.updateValue(record)
  }
  handleDelete = (record) => {
    const { id } = record
    deleteGroup({}, { id }, (res) => {
      message.success(`删除 ${id} 成功`)
      this.getList()
    })
  }
  rulesUpdate = rule => new Promise((resolve) => {
    const { id, ...data } = rule
    if (id) {
      updateGroup(data, { id }, (res) => {
        this.getList()
        resolve(true)
      })
      return
    }
    addGroup(data, (res) => {
      resolve(true)
      this.getList()
    })
  })
  onRef(component) {
    this.createEditGroup = component
  }
  render() {
    const { dataSource } = this.state;
    const columns = [
      { title: '编号', align: 'center', dataIndex: 'id' },
      {
        title: '组名',
        align: 'center',
        dataIndex: 'name',
        ...this.getColumnSearchProps('name'),
        filterDropdownVisible: this.state.filterItem.name,
      },
      {
        title: '成员',
        align: 'center',
        dataIndex: 'user',
        ...this.getColumnSearchProps('user'),
        filterDropdownVisible: this.state.filterItem.user,
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
              onConfirm={() => this.handleDelete(record)}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">删除</a>
            </Popconfirm>
          </span>
        ),
      },
    ];
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <CreateEditGroup onRef={c => this.onRef(c)} onSubmit={this.rulesUpdate} />
        <Table dataSource={dataSource} columns={columns} rowKey="id" />
      </div>
    )
  }
}
