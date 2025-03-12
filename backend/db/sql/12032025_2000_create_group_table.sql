CREATE TABLE public.group (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_by UUID NOT NULL,
    created_at timestamp NOT NULL
);

ALTER TABLE public.group
ADD CONSTRAINT fk_group_user
FOREIGN KEY (created_by)
REFERENCES public.user (id);

CREATE TABLE public.group_member (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at timestamp NOT NULL
);

ALTER TABLE public.group_member
ADD CONSTRAINT fk_group_member_group
FOREIGN KEY (group_id)
REFERENCES public.group (id);

ALTER TABLE public.group_member
ADD CONSTRAINT fk_group_member_user
FOREIGN KEY (user_id)
REFERENCES public.user (id);