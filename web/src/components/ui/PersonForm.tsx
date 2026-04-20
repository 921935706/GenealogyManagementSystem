import { useForm, zodResolver } from '@mantine/form'
import { TextInput, Button, Stack, Select, Checkbox, Group, Textarea, Grid, Title } from '@mantine/core'
import { z } from 'zod'
import { Person } from '@/types/api'

const personSchema = z.object({
  name: z.string().min(2, '姓名至少2个字符'),
  firstName: z.string().optional(),
  lastName: z.string().optional(),
  generationName: z.string().optional(),
  alias: z.string().optional(),
  englishName: z.string().optional(),
  gender: z.enum(['male', 'female'], {
    errorMap: () => ({ message: '请选择性别' }),
  }),
  isAlive: z.boolean().default(true),
  birthDate: z.string()
    .optional()
    .refine((val) => !val || /^(公元前)?\d{4}-\d{2}-\d{2}$/.test(val), {
      message: '日期格式应为 YYYY-MM-DD 或 公元前YYYY-MM-DD'
    }),
  birthLunar: z.string().optional(),
  birthPlace: z.string().optional(),
  birthPlaceLongitude: z.string().optional(),
  birthPlaceLatitude: z.string().optional(),
  deathDate: z.string()
    .optional()
    .refine((val) => !val || /^(公元前)?\d{4}-\d{2}-\d{2}$/.test(val), {
      message: '日期格式应为 YYYY-MM-DD 或 公元前YYYY-MM-DD'
    }),
  deathLunar: z.string().optional(),
  deathPlace: z.string().optional(),
  deathPlaceLongitude: z.string().optional(),
  deathPlaceLatitude: z.string().optional(),
  deathCause: z.string().optional(),
  generation: z.string().optional(),
  occupation: z.string().optional(),
  education: z.string().optional(),
  biography: z.string().optional(),
  avatarUrl: z.string().url('头像链接格式不正确').optional(),
  coverPhotoUrl: z.string().url('封面照片链接格式不正确').optional(),
  isPublic: z.boolean().default(true),
  verificationStatus: z.string().optional(),
  createdBy: z.string().optional(),
  email: z.string().email('邮箱格式不正确').optional(),
})

type PersonFormData = z.infer<typeof personSchema>

interface PersonFormProps {
  onSubmit: (data: PersonFormData) => void
  initialValues?: Partial<Person>
}

export function PersonForm({ onSubmit, initialValues }: PersonFormProps) {
  const form = useForm({
    initialValues: {
      name: '',
      firstName: '',
      lastName: '',
      generationName: '',
      alias: '',
      englishName: '',
      gender: 'male' as 'male' | 'female',
      isAlive: true,
      birthDate: '',
      birthLunar: '',
      birthPlace: '',
      birthPlaceLongitude: '',
      birthPlaceLatitude: '',
      deathDate: '',
      deathLunar: '',
      deathPlace: '',
      deathPlaceLongitude: '',
      deathPlaceLatitude: '',
      deathCause: '',
      generation: '',
      occupation: '',
      education: '',
      biography: '',
      avatarUrl: '',
      coverPhotoUrl: '',
      isPublic: true,
      verificationStatus: '',
      createdBy: '',
      email: '',
      ...initialValues,
    },
    validate: zodResolver(personSchema),
  })

  return (
    <form onSubmit={form.onSubmit(onSubmit)}>
      <Stack gap="md">
        <Grid>
          <Grid.Col span={4}>
            <TextInput
              label="姓名"
              placeholder="请输入姓名"
              required
              {...form.getInputProps('name')}
            />
          </Grid.Col>
          <Grid.Col span={4}>
            <TextInput
              label="英文名"
              placeholder="请输入英文名"
              {...form.getInputProps('englishName')}
            />
          </Grid.Col>
          <Grid.Col span={4}>
            <Select
              label="性别"
              placeholder="请选择性别"
              data={[
                { value: 'male', label: '男' },
                { value: 'female', label: '女' },
              ]}
              required
              {...form.getInputProps('gender')}
            />
          </Grid.Col>
        </Grid>

        <Grid>
          <Grid.Col span={3}>
            <TextInput
              label="姓氏"
              placeholder="请输入姓氏"
              {...form.getInputProps('lastName')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="名字"
              placeholder="请输入名字"
              {...form.getInputProps('firstName')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="字辈"
              placeholder="请输入字辈"
              {...form.getInputProps('generationName')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="辈分"
              placeholder="请输入辈分"
              {...form.getInputProps('generation')}
            />
          </Grid.Col>
        </Grid>

        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="别名/号"
              placeholder="请输入别名或号"
              {...form.getInputProps('alias')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="创建者"
              placeholder="请输入创建者"
              {...form.getInputProps('createdBy')}
            />
          </Grid.Col>
        </Grid>

        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="邮箱"
              placeholder="请输入邮箱"
              {...form.getInputProps('email')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <Select
              label="验证状态"
              placeholder="请选择验证状态"
              data={[
                { value: 'pending', label: '待验证' },
                { value: 'verified', label: '已验证' },
                { value: 'rejected', label: '已拒绝' },
              ]}
              {...form.getInputProps('verificationStatus')}
            />
          </Grid.Col>
        </Grid>

        <Group grow>
          <Checkbox
            label="是否在世"
            {...form.getInputProps('isAlive', { type: 'checkbox' })}
          />
          <Checkbox
            label="是否公开"
            {...form.getInputProps('isPublic', { type: 'checkbox' })}
          />
          <Select
            label="验证状态"
            placeholder="请选择验证状态"
            data={[
              { value: 'pending', label: '待验证' },
              { value: 'verified', label: '已验证' },
              { value: 'rejected', label: '已拒绝' },
            ]}
            {...form.getInputProps('verificationStatus')}
          />
        </Group>

        <Title order={4}>出生信息</Title>
        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="出生日期"
              placeholder="请输入出生日期 (YYYY-MM-DD 或 公元前YYYY-MM-DD)"
              {...form.getInputProps('birthDate')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="农历生日"
              placeholder="请输入农历生日"
              {...form.getInputProps('birthLunar')}
            />
          </Grid.Col>
        </Grid>

        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="出生地点"
              placeholder="请输入出生地点"
              {...form.getInputProps('birthPlace')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="经度"
              placeholder="经度"
              {...form.getInputProps('birthPlaceLongitude')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="纬度"
              placeholder="纬度"
              {...form.getInputProps('birthPlaceLatitude')}
            />
          </Grid.Col>
        </Grid>

        <Title order={4}>逝世信息</Title>
        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="逝世日期"
              placeholder="请输入逝世日期 (YYYY-MM-DD 或 公元前YYYY-MM-DD)"
              {...form.getInputProps('deathDate')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="农历逝世日期"
              placeholder="请输入农历逝世日期"
              {...form.getInputProps('deathLunar')}
            />
          </Grid.Col>
        </Grid>

        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="逝世地点"
              placeholder="请输入逝世地点"
              {...form.getInputProps('deathPlace')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="经度"
              placeholder="经度"
              {...form.getInputProps('deathPlaceLongitude')}
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="纬度"
              placeholder="纬度"
              {...form.getInputProps('deathPlaceLatitude')}
            />
          </Grid.Col>
        </Grid>

        <TextInput
          label="逝世原因"
          placeholder="请输入逝世原因"
          {...form.getInputProps('deathCause')}
        />

        <Title order={4}>职业与教育</Title>
        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="职业"
              placeholder="请输入职业"
              {...form.getInputProps('occupation')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="教育程度"
              placeholder="请输入教育程度"
              {...form.getInputProps('education')}
            />
          </Grid.Col>
        </Grid>

        <Title order={4}>媒体信息</Title>
        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="头像链接"
              placeholder="请输入头像链接"
              {...form.getInputProps('avatarUrl')}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="封面照片链接"
              placeholder="请输入封面照片链接"
              {...form.getInputProps('coverPhotoUrl')}
            />
          </Grid.Col>
        </Grid>

        <Textarea
          label="生平简介"
          placeholder="请输入生平简介"
          rows={3}
          {...form.getInputProps('biography')}
        />
        
        <Button type="submit" fullWidth loading={false}>
          保存
        </Button>
      </Stack>
    </form>
  )
}