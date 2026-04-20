import { useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { 
  Card, 
  Text, 
  Group, 
  Badge, 
  Button, 
  Stack, 
  Title, 
  Loader,
  Alert,
  Grid
} from '@mantine/core'
import { IconArrowLeft, IconEdit } from '@tabler/icons-react'
import { usePersonStore } from '@/stores/personStore'

export function PersonDetail() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  
  const { 
    currentPerson, 
    loading, 
    error, 
    fetchPersonById 
  } = usePersonStore()

  useEffect(() => {
    if (id) {
      fetchPersonById(parseInt(id))
    }
  }, [id])

  if (loading) {
    return (
      <Group justify="center">
        <Loader />
      </Group>
    )
  }

  if (error) {
    return (
      <Alert color="red" title="错误">
        {error}
      </Alert>
    )
  }

  if (!currentPerson) {
    return (
      <Alert color="yellow" title="提示">
        未找到该人员信息
      </Alert>
    )
  }

  return (
    <Stack gap="md">
      <Group>
        <Button 
          variant="subtle" 
          leftSection={<IconArrowLeft size={16} />}
          onClick={() => navigate('/persons')}
        >
          返回列表
        </Button>
        <Button 
          leftSection={<IconEdit size={16} />}
          onClick={() => navigate(`/persons/edit/${currentPerson.id}`)}
        >
          编辑
        </Button>
      </Group>

      <Card shadow="sm" padding="lg">
        <Title order={2} mb="md">{currentPerson.name}</Title>
        
        <Grid>
          <Grid.Col span={6}>
            <Stack gap="sm">
              <Group>
                <Text fw={500}>性别:</Text>
                <Badge color={currentPerson.gender === 'male' ? 'blue' : 'pink'}>
                  {currentPerson.gender === 'male' ? '男' : '女'}
                </Badge>
              </Group>
              
              <Group>
                <Text fw={500}>状态:</Text>
                <Badge color={currentPerson.isAlive ? 'green' : 'red'}>
                  {currentPerson.isAlive ? '在世' : '已故'}
                </Badge>
              </Group>
              
              <Group>
                <Text fw={500}>公开状态:</Text>
                <Badge color={currentPerson.isPublic ? 'green' : 'gray'}>
                  {currentPerson.isPublic ? '公开' : '私有'}
                </Badge>
              </Group>
              
              {currentPerson.birthDate && (
                <Group>
                  <Text fw={500}>出生日期:</Text>
                  <Text>{currentPerson.birthDate}</Text>
                </Group>
              )}
              
              {currentPerson.email && (
                <Group>
                  <Text fw={500}>邮箱:</Text>
                  <Text>{currentPerson.email}</Text>
                </Group>
              )}
            </Stack>
          </Grid.Col>
          
          <Grid.Col span={6}>
            <Stack gap="sm">
              {currentPerson.occupation && (
                <Group>
                  <Text fw={500}>职业:</Text>
                  <Text>{currentPerson.occupation}</Text>
                </Group>
              )}
              
              {currentPerson.education && (
                <Group>
                  <Text fw={500}>教育程度:</Text>
                  <Text>{currentPerson.education}</Text>
                </Group>
              )}
              
              {currentPerson.generation && (
                <Group>
                  <Text fw={500}>世代:</Text>
                  <Text>{currentPerson.generation}</Text>
                </Group>
              )}
              
              {currentPerson.birthPlace && (
                <Group>
                  <Text fw={500}>出生地:</Text>
                  <Text>{currentPerson.birthPlace}</Text>
                </Group>
              )}
            </Stack>
          </Grid.Col>
        </Grid>
        
        {currentPerson.biography && (
          <Card.Section mt="md" p="md" bg="gray.0">
            <Text fw={500} mb="sm">生平简介:</Text>
            <Text>{currentPerson.biography}</Text>
          </Card.Section>
        )}
      </Card>
    </Stack>
  )
}