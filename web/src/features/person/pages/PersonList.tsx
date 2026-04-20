import { useState, useEffect } from 'react'
import { 
  Table, 
  Button, 
  Group, 
  TextInput, 
  ActionIcon, 
  Badge, 
  Modal,
  Title,
  Stack,
  Pagination,
  Loader,
  Alert
} from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import { IconEdit, IconTrash, IconSearch, IconPlus } from '@tabler/icons-react'
import { usePersonStore } from '@/stores/personStore'
import { PersonForm } from '@/components/ui/PersonForm'
import { Person } from '@/types/api'

export function PersonList() {
  const [searchTerm, setSearchTerm] = useState('')
  const [currentPage, setCurrentPage] = useState(1)
  const [editingPerson, setEditingPerson] = useState<Person | null>(null)
  const [modalOpened, { open: openModal, close: closeModal }] = useDisclosure(false)
  
  const { 
    persons, 
    loading, 
    error, 
    total, 
    fetchPersons, 
    deletePerson, 
    searchPersons,
    createPerson,
    updatePerson
  } = usePersonStore()

  const pageSize = 10

  useEffect(() => {
    if (searchTerm) {
      searchPersons(searchTerm, currentPage, pageSize)
    } else {
      fetchPersons(currentPage, pageSize)
    }
  }, [currentPage, searchTerm])

  const handleSearch = () => {
    setCurrentPage(1)
    if (searchTerm) {
      searchPersons(searchTerm, 1, pageSize)
    } else {
      fetchPersons(1, pageSize)
    }
  }

  const handleCreate = () => {
    setEditingPerson(null)
    openModal()
  }

  const handleEdit = (person: Person) => {
    setEditingPerson(person)
    openModal()
  }

  const handleDelete = async (id: number) => {
    if (window.confirm('确定要删除这个人员吗？')) {
      const success = await deletePerson(id)
      if (success) {
        if (searchTerm) {
          searchPersons(searchTerm, currentPage, pageSize)
        } else {
          fetchPersons(currentPage, pageSize)
        }
      }
    }
  }

  const handleFormSubmit = async (data: Person) => {
    const success = editingPerson 
      ? await updatePerson({ ...editingPerson, ...data })
      : await createPerson(data)
    
    if (success) {
      closeModal()
      setEditingPerson(null)
      if (searchTerm) {
        searchPersons(searchTerm, currentPage, pageSize)
      } else {
        fetchPersons(currentPage, pageSize)
      }
    }
  }

  const rows = persons.map((person) => (
    <Table.Tr key={person.id}>
      <Table.Td>{person.name}</Table.Td>
      <Table.Td>
        <Badge color={person.gender === 'male' ? 'blue' : 'pink'}>
          {person.gender === 'male' ? '男' : '女'}
        </Badge>
      </Table.Td>
      <Table.Td>
        <Badge color={person.isAlive ? 'green' : 'red'}>
          {person.isAlive ? '在世' : '已故'}
        </Badge>
      </Table.Td>
      <Table.Td>{person.birthDate || '-'}</Table.Td>
      <Table.Td>{person.occupation || '-'}</Table.Td>
      <Table.Td>
        <Group gap="xs">
          <ActionIcon 
            variant="subtle" 
            color="blue" 
            onClick={() => handleEdit(person)}
          >
            <IconEdit size={16} />
          </ActionIcon>
          <ActionIcon 
            variant="subtle" 
            color="red" 
            onClick={() => handleDelete(person.id!)}
          >
            <IconTrash size={16} />
          </ActionIcon>
        </Group>
      </Table.Td>
    </Table.Tr>
  ))

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Title order={2}>人员管理</Title>
        <Button 
          leftSection={<IconPlus size={16} />} 
          onClick={handleCreate}
        >
          新建人员
        </Button>
      </Group>

      <Group>
        <TextInput
          placeholder="搜索人员姓名..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
          rightSection={
            <ActionIcon onClick={handleSearch}>
              <IconSearch size={16} />
            </ActionIcon>
          }
        />
      </Group>

      {error && (
        <Alert color="red" title="错误">
          {error}
        </Alert>
      )}

      {loading ? (
        <Group justify="center">
          <Loader />
        </Group>
      ) : (
        <>
          <Table highlightOnHover>
            <Table.Thead>
              <Table.Tr>
                <Table.Th>姓名</Table.Th>
                <Table.Th>性别</Table.Th>
                <Table.Th>状态</Table.Th>
                <Table.Th>出生日期</Table.Th>
                <Table.Th>职业</Table.Th>
                <Table.Th>操作</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {rows.length > 0 ? rows : (
                <Table.Tr>
                  <Table.Td colSpan={6} style={{ textAlign: 'center' }}>
                    暂无数据
                  </Table.Td>
                </Table.Tr>
              )}
            </Table.Tbody>
          </Table>

          {total > 0 && (
            <Group justify="center">
              <Pagination
                value={currentPage}
                onChange={setCurrentPage}
                total={Math.ceil(total / pageSize)}
              />
            </Group>
          )}
        </>
      )}

      <Modal 
        opened={modalOpened} 
        onClose={() => {
          closeModal()
          setEditingPerson(null)
        }}
        title={editingPerson ? '编辑人员' : '新建人员'}
        size="80%"
        centered
        styles={{
          content: {
            maxWidth: '1200px',
            margin: '0 auto'
          },
          body: {
            maxHeight: '80vh',
            overflowY: 'auto'
          }
        }}
      >
        <PersonForm
          onSubmit={handleFormSubmit}
          initialValues={editingPerson || undefined}
        />
      </Modal>
    </Stack>
  )
}