# 族谱管理系统 - web 页面设计文档

## 🎯 核心原则

1. Mantine 优先：UI 组件必须从 @mantine/core 导入，禁止手写 CSS 实现组件功能（如 Modal、Drawer、Popover 等）
2. Tailwind 辅助：布局和间距可以使用 Tailwind CSS，但组件样式优先使用 Mantine 的 styles 或 className
3. TypeScript 强制：所有组件、Props、状态必须有明确的类型定义
4. 移动端优先：使用 Mantine 的 useMatches 或 createStyles 实现响应式

## 📦 技术栈

- 框架: React 18+ (Vite 构建)
- UI 库: Mantine v7 (最新版)
- 样式: Mantine + Tailwind CSS 3 (仅用于快速布局)
- 路由: React Router v6 (如需要)
- 状态管理: Zustand 或 TanStack Query (服务端状态)
- 图标: @tabler/icons-react

## 🏗️ 推荐项目结构

```
src/
├── components/          # 通用组件（与业务无关）
│   ├── ui/             # Mantine 二次封装的组件
│   │   ├── DataTable.tsx
│   │   └── ModalForm.tsx
│   └── layout/         # 布局组件
│       ├── AppShell.tsx
│       └── Navbar.tsx
├── features/           # 功能模块（按业务划分）
│   ├── auth/
│   │   ├── components/ # 模块私有组件
│   │   ├── hooks/      # 模块私有 hooks
│   │   ├── types/      # 模块类型定义
│   │   └── pages/      # 页面组件
│   └── dashboard/
├── hooks/              # 全局自定义 hooks
├── lib/                # 第三方配置（如 Mantine 主题、axios）
├── stores/             # Zustand stores
├── types/              # 全局类型定义
├── utils/              # 工具函数
├── App.tsx
└── main.tsx
```
## 🎨 Mantine 核心使用规范

### 1. 主题配置 (必须)

在 src/lib/mantine.ts 中集中定义主题：
```typescript
import { createTheme, MantineThemeOverride } from '@mantine/core';

export const theme: MantineThemeOverride = createTheme({
  primaryColor: 'blue',
  fontFamily: 'Inter, sans-serif',
  defaultRadius: 'md',
  components: {
    Button: {
      defaultProps: {
        variant: 'filled',
      },
    },
  },
});
```

### 2. 样式编写规范

#### 规则 1：优先使用 Mantine 内置样式系统

```
✅ 正确
<Stack gap="md" p="lg" bg="gray.0">
  <Title order={3}>标题</Title>
</Stack>
```

❌ 错误

```
<div style={{ display: 'flex', gap: '1rem' }}>...</div>
```

#### 规则 2：需要自定义样式时使用 createStyles

import { createStyles } from '@mantine/core';

```
const useStyles = createStyles((theme) => ({
  card: {
    backgroundColor: theme.colors.gray[0],
    transition: 'transform 0.2s',
    '&:hover': {
      transform: 'scale(1.01)',
    },
  },
}));

function MyComponent() {
  const { classes } = useStyles();
  return <Card className={classes.card}>...</Card>;
}
```

#### 规则 3：动态样式使用 style prop 或 cx 合并
```
<Button
  style={{ display: isVisible ? 'block' : 'none' }}
  className={cx(classes.button, { [classes.active]: isActive })}
>
```
### 3. 表单处理规范

#### 规则 必须使用 Mantine Form + Zod 组合

```
import { useForm, zodResolver } from '@mantine/form';
import { z } from 'zod';

const schema = z.object({
  email: z.string().email('邮箱格式不正确'),
  password: z.string().min(6, '密码至少6位'),
});

function LoginForm() {
  const form = useForm({
    initialValues: { email: '', password: '' },
    validate: zodResolver(schema),
  });

  return (
    <form onSubmit={form.onSubmit((values) => console.log(values))}>
      <TextInput
        label="邮箱"
        {...form.getInputProps('email')}
      />
      <PasswordInput
        label="密码"
        {...form.getInputProps('password')}
      />
      <Button type="submit">登录</Button>
    </form>
  );
}
```

### 4. 数据获取规范

#### 规则 必须使用 TanStack Query (推荐) 或 useEffect + fetch

```
import { useQuery } from '@tanstack/react-query';

function UserList() {
  const { data, isLoading, error } = useQuery({
    queryKey: ['users'],
    queryFn: () => fetch('/api/users').then(res => res.json()),
  });

  if (isLoading) return <Loader />;
  if (error) return <Alert color="red">加载失败</Alert>;

  return (
    <SimpleGrid cols={3}>
      {data.map(user => <UserCard key={user.id} {...user} />)}
    </SimpleGrid>
  );
}
```

## 🧩 常用组件使用示例

### 1. 页面布局
```
import { AppShell, Burger, NavLink } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';

function Layout({ children }) {
  const [opened, { toggle }] = useDisclosure();

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: 'sm', collapsed: { mobile: !opened } }}
      padding="md"
    >
      <AppShell.Header>
        <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
      </AppShell.Header>
      <AppShell.Navbar>
        <NavLink label="首页" href="/" />
        <NavLink label="设置" href="/settings" />
      </AppShell.Navbar>
      <AppShell.Main>{children}</AppShell.Main>
    </AppShell>
  );
}
```
### 2. 数据表格

```
import { Table, Badge, ActionIcon } from '@mantine/core';
import { IconEdit, IconTrash } from '@tabler/icons-react';

function DataTable({ data }) {
  const rows = data.map((row) => (
    <Table.Tr key={row.id}>
      <Table.Td>{row.name}</Table.Td>
      <Table.Td>
        <Badge color={row.active ? 'green' : 'red'}>
          {row.active ? '启用' : '禁用'}
        </Badge>
      </Table.Td>
      <Table.Td>
        <ActionIcon variant="subtle" onClick={() => edit(row.id)}>
          <IconEdit size={16} />
        </ActionIcon>
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Table highlightOnHover>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>名称</Table.Th>
          <Table.Th>状态</Table.Th>
          <Table.Th>操作</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>{rows}</Table.Tbody>
    </Table>
  );
}
```

### 3. 模态框表单

```
import { Modal, Button, TextInput } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';

function CreateModal() {
  const [opened, { open, close }] = useDisclosure(false);

  return (
    <>
      <Button onClick={open}>新建</Button>
      <Modal
        opened={opened}
        onClose={close}
        title="创建项目"
        size="lg"
      >
        <TextInput label="名称" placeholder="输入项目名称" />
        {/* 表单内容 */}
      </Modal>
    </>
  );
}
```

## 📱 响应式设计规范

### 使用 useMatches Hook

```
import { useMatches } from '@mantine/core';

function ResponsiveGrid() {
  const cols = useMatches({
    base: 1,
    xs: 2,
    sm: 3,
    md: 4,
    lg: 5,
  });

  return <SimpleGrid cols={cols}>...</SimpleGrid>;
}
```

## 🚫 禁止行为清单

1. 禁止使用原生 HTML 弹窗（alert/confirm），必须用 Mantine Modal
2. 禁止手写 CSS 实现布局（如 flex/grid），必须用 Mantine 的 Group/Stack/Grid
3. 禁止使用 any 类型
4. 禁止在组件内直接写 fetch，必须封装成 hook 或使用 TanStack Query
5. 禁止使用内联 style 实现主题相关的样式（应使用 Mantine 主题系统）

## ✅ 代码检查清单

生成代码前，请确认：

- [ ] 所有导入来自正确的包（@mantine/core, @mantine/hooks）
- [ ] 使用了 Mantine 的组件而非原生 HTML（如 <Button> 而非 <button>）
- [ ] 表单使用了 useForm 和 Zod 验证
- [ ] 状态管理使用了 Zustand 或 TanStack Query
- [ ] 响应式使用了 useMatches 或 Mantine 的响应式 props
- [ ] 没有出现 any 类型
- [ ] 组件有明确的 Props 类型定义

## 📚 快速参考

| 需求 | 使用组件/Hook |
|------|--------------|
| 布局 | AppShell, Container, Grid, Stack, Group |
| 表单 | useForm + TextInput, Select, Checkbox |
| 数据展示 | Table, Card, Badge, Avatar |
| 反馈 | Modal, Drawer, Notification, Tooltip |
| 导航 | Tabs, NavLink, Breadcrumbs |
| 状态 | useDisclosure, useToggle, useLocalStorage |

当 AI 不确定使用哪个 Mantine 组件时，应查阅 Mantine 官方文档后给出答案。